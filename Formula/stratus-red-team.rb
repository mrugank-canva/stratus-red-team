# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class StratusRedTeam < Formula
  desc ""
  homepage "https://stratus-red-team.cloud"
  version "2.5.2"
  license "Apache-2.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.2/stratus-red-team_2.5.2_Darwin_arm64.tar.gz"
      sha256 "1a4d31309458faa6d8495c565be8c4f09f1e7b31b544480f6d5600c457f332ac"

      def install
        bin.install "stratus"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.2/stratus-red-team_2.5.2_Darwin_x86_64.tar.gz"
      sha256 "6eb2ec769e89bb479ca013bdb5c4bc354ba299df8f5e4dc4212bf54c1c822fe4"

      def install
        bin.install "stratus"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.2/stratus-red-team_2.5.2_Linux_x86_64.tar.gz"
      sha256 "c5e158c0d498ca77c90f35610f4694c417e5485a75d9afdaedce2bb5d0290c6d"

      def install
        bin.install "stratus"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/DataDog/stratus-red-team/releases/download/v2.5.2/stratus-red-team_2.5.2_Linux_arm64.tar.gz"
      sha256 "79cac8b619f2293355d8382657bc41ae9fe8b44ad04d6dba4bf1d0da0504e5c9"

      def install
        bin.install "stratus"
      end
    end
  end
end
