package semantics_test
import("net";"net/http";"net/http/httptest";"testing";echo "github.com/labstack/echo/v5")
func TestPublicScopedAddressSemantics(t *testing.T){
 for _,tc:=range []struct{name,remote,xff,trusted,wantIP,wantScheme string}{
 {"uppercase-expanded","[FE80:0000:0000:0000:0000:0000:0000:000A%ETH0]:8080","198.51.100.7","fe80::/64","198.51.100.7","https"},
 {"compressed","[fe80::a%eth0]:8080","198.51.100.7","fe80::/64","198.51.100.7","https"},
 {"noncanonical-expanded","[FE80:0:0:0:0:0:0:A%5]:8080","198.51.100.7","fe80::/64","198.51.100.7","https"},
 {"IPv4-mapped","[::ffff:127.0.0.1%eth0]:8080","198.51.100.7","127.0.0.0/8","198.51.100.7","https"},
 {"IPv4-mapped-hex","[::FFFF:7F00:0001%eth0]:8080","198.51.100.7","127.0.0.0/8","198.51.100.7","https"},
 {"noncanonical-untrusted-header","127.0.0.1:8080","2001:DB8:0:0:0:0:0:A%uplink","127.0.0.0/8","2001:db8::a","https"},
 {"invalid-IPv4-zone","127.0.0.1%garbage:8080","198.51.100.7","127.0.0.0/8","127.0.0.1%garbage","http"},
 {"malformed-IPv6","[fe80:::1%eth0]:8080","198.51.100.7","fe80::/64","fe80:::1%eth0","http"},
 {"empty-zone","[fe80::1%]:8080","198.51.100.7","fe80::/64","fe80::1%","http"},
 {"malformed-IPv4-mapped","[::ffff:127.0.0.999%eth0]:8080","198.51.100.7","127.0.0.0/8","::ffff:127.0.0.999%eth0","http"},
 }{t.Run(tc.name,func(t *testing.T){
  _,network,err:=net.ParseCIDR(tc.trusted);if err!=nil{t.Fatal(err)}
  opts:=[]echo.TrustOption{echo.TrustLoopback(false),echo.TrustLinkLocal(false),echo.TrustPrivateNet(false),echo.TrustIPRange(network)}
  e:=echo.New();e.IPExtractor=echo.ExtractIPFromXFFHeader(opts...);e.SchemeExtractor=echo.ExtractSchemeFromHeaders(opts...)
  req:=httptest.NewRequest(http.MethodGet,"http://example.test/",nil);req.RemoteAddr=tc.remote;req.Header.Set(echo.HeaderXForwardedFor,tc.xff);req.Header.Set(echo.HeaderXForwardedProto,"https")
  c:=e.NewContext(req,httptest.NewRecorder());if got:=c.RealIP();got!=tc.wantIP{t.Errorf("RealIP got=%q want=%q",got,tc.wantIP)};if got:=c.Scheme();got!=tc.wantScheme{t.Errorf("Scheme got=%q want=%q",got,tc.wantScheme)}
 })}
}
