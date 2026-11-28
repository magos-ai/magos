// swift-tools-version: 6.4
import PackageDescription

let package = Package(
  name: "Magos",
  platforms: [.macOS(.v27)],
  targets: [
    .executableTarget(name: "Magos")
  ]
)
