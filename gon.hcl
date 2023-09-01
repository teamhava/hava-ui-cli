# gon.hcl
#
# The path follows a pattern
# ./dist/BUILD-ID_TARGET/BINARY-NAME
source = ["./dist/hava-macos_darwin_amd64_v1/hava", "./dist/hava-macos_darwin_arm64/hava"]
bundle_id = "io.hava.cli"

// set using env vars
// apple_id {
//   password = "@env:AC_PASSWORD"
// }

sign {
  application_identity = "Developer ID Application: Thomas Winsnes"
}

zip {
  output_path = "hava.zip"
}