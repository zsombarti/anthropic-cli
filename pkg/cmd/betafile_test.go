// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/anthropic-cli/internal/mocktest"
)

func TestBetaFilesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"beta:files", "list",
		"--after-id", "after_id",
		"--before-id", "before_id",
		"--limit", "1",
		"--beta", "string",
	)
}

func TestBetaFilesDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"beta:files", "delete",
		"--file-id", "file_id",
		"--beta", "string",
	)
}

func TestBetaFilesDownload(t *testing.T) {
	t.Skip("Prism doesn't support application/binary responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"beta:files", "download",
		"--file-id", "file_id",
		"--beta", "string",
	)
}

func TestBetaFilesRetrieveMetadata(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"beta:files", "retrieve-metadata",
		"--file-id", "file_id",
		"--beta", "string",
	)
}

func TestBetaFilesUpload(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"beta:files", "upload",
		"--file", "",
		"--beta", "string",
	)
}
