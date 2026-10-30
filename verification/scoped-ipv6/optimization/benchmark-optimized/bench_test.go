package scopedip_test
import("net/http";"testing";echo "github.com/labstack/echo/v5")
var result string
func BenchmarkXFFIPv4(b *testing.B){ f:=echo.ExtractIPFromXFFHeader();r:=&http.Request{RemoteAddr:"127.0.0.1:8080",Header:http.Header{echo.HeaderXForwardedFor:[]string{"198.51.100.7"}}};b.ReportAllocs();for b.Loop(){result=f(r)}}
func BenchmarkXFFScopedIPv6(b *testing.B){ f:=echo.ExtractIPFromXFFHeader();r:=&http.Request{RemoteAddr:"[fe80::1%eth0]:8080",Header:http.Header{echo.HeaderXForwardedFor:[]string{"198.51.100.7"}}};b.ReportAllocs();for b.Loop(){result=f(r)}}
func BenchmarkSchemeScopedIPv6(b *testing.B){ f:=echo.ExtractSchemeFromHeaders();r:=&http.Request{RemoteAddr:"[fe80::1%eth0]:8080",Header:http.Header{echo.HeaderXForwardedProto:[]string{"https"}}};b.ReportAllocs();for b.Loop(){result=f(r)}}
