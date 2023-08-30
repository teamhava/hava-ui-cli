# gon.hcl
#
# The path follows a pattern
# ./dist/BUILD-ID_TARGET/BINARY-NAME
source = ["./dist/hava-macos_darwin_amd64_v1/hava", "./dist/hava-macos_darwin_arm64/hava"]
bundle_id = "io.hava.cli"

apple_id {
  username = "@env:APPLE_USERNAME"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Installer: Thomas Winsnes (WB8X3342BR)"
}
