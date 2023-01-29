package musik

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWacipher(t *testing.T) {
	portdangdut := Dangdut()
	require.Equal(t, ":80", portdangdut)

	portkoplo := Koplo()
	require.Equal(t, ":3000", portkoplo)

	porttarling := Tarling()
	require.Equal(t, "80", porttarling)

}
