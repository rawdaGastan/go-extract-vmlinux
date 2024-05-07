package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKernels(t *testing.T) {
	kernels := []string{"gzip", "xz", "lzma", "lzo", "lz4", "zstd", "bzip2"}

	for _, kernel := range kernels {
		kernelImagePath := fmt.Sprintf("kernels/bzImage-%s", kernel)
		outputKernelPath := kernelImagePath + "-decompressed"

		// compressed kernel
		require.Error(t, IsValidELFKernel(kernelImagePath))

		// decompress kernel
		require.NoError(t, TryDecompressKernel(kernelImagePath, outputKernelPath))

		// compressed kernel
		require.NoError(t, IsValidELFKernel(outputKernelPath))
	}
}
