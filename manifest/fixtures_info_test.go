package manifest

import "github.com/opencontainers/go-digest"

const (
	// TestV2S2ManifestDigest is the Docker manifest digest of "v2s2.manifest.json"
	TestDockerV2S2ManifestDigest = digest.Digest("sha256:20bf21ed457b390829cdbeec8795a7bea1626991fda603e0d01b4e7f60427e55")
	// TestV2S1ManifestDigest is the Docker manifest digest of "v2s1.manifest.json"
	TestDockerV2S1ManifestDigest = digest.Digest("sha256:7364fea9d84ee548ab67d4c46c6006289800c98de3fbf8c0a97138dfcc23f000")
	// TestV2S1UnsignedManifestDigest is the Docker manifest digest of "v2s1unsigned.manifest.json"
	TestDockerV2S1UnsignedManifestDigest = digest.Digest("sha256:7364fea9d84ee548ab67d4c46c6006289800c98de3fbf8c0a97138dfcc23f000")
)
