package testdata

import (
	"testing"
)

// Buf is a ...
type Buf []byte

// Append ...
func (*Buf) Append([]byte) {}

func (Buf) Reset() {}

func (Buf) Len() int { return 0 }

// DefaultBuf is a ...
var DefaultBuf Buf

func Example() {} // OK because is package-level.

func Example_goodSuffix() // OK because refers to suffix annotation.

func Example_BadSuffix() // ERROR "Example_BadSuffix has malformed example suffix: BadSuffix"

func ExampleBuf() // OK because refers to known top-level type.

func ExampleBuf_Append() {} // OK because refers to known method.

func ExampleBuf_Clear() {} // ERROR "ExampleBuf_Clear refers to unknown field or method: Buf.Clear"

func ExampleBuf_suffix() {} // OK because refers to suffix annotation.

func ExampleBuf_Append_Bad() {} // ERROR "ExampleBuf_Append_Bad has malformed example suffix: Bad"

func ExampleBuf_Append_suffix() {} // OK because refers to known method with valid suffix.

func ExampleDefaultBuf() {} // OK because refers to top-level identifier.

func ExampleBuf_Reset() bool { return true } // ERROR "ExampleBuf_Reset should return nothing"

func ExampleBuf_Len(i int) {} // ERROR "ExampleBuf_Len should be niladic"

// "Puffer" is German for "Buffer".

func ExamplePuffer() // ERROR "ExamplePuffer refers to unknown identifier: Puffer"

func ExamplePuffer_Append() // ERROR "ExamplePuffer_Append refers to unknown identifier: Puffer"

func ExamplePuffer_suffix() // ERROR "ExamplePuffer_suffix refers to unknown identifier: Puffer"

func nonTest() {} // OK because it doesn't start with "Test".

func (Buf) TesthasReceiver() {} // OK because it has a receiver.

func TestOKSuffix(*testing.T) {} // OK because first char after "Test" is Uppercase.

func TestÜnicodeWorks(*testing.T) {} // OK because the first char after "Test" is Uppercase.

func TestbadSuffix(*testing.T) {} // ERROR "first letter after 'Test' must not be lowercase"

func TestemptyImportBadSuffix(*T) {} // ERROR "first letter after 'Test' must not be lowercase"

func Test(*testing.T) {} // OK "Test" on its own is considered a test.

func Testify() {} // OK because it takes no parameters.

func TesttooManyParams(*testing.T, string) {} // OK because it takes too many parameters.

func TesttooManyNames(a, b *testing.T) {} // OK because it takes too many names.

func TestnoTParam(string) {} // OK because it doesn't take a *testing.T

func BenchmarkbadSuffix(*testing.B) {} // ERROR "first letter after 'Benchmark' must not be lowercase"
