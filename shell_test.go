package shellgo_test

import (
	"fmt"
	sgo "github.com/ElPotato/shellgo"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const inputDir = "./test_inputs/"

func Test_parseOutputValidity(t *testing.T) {
	inputs := map[string]string{
		inputDir + "0_test_input.txt": "4883ec2048896c2418488d6c241848c74424280000000048c704240500000048c74424080a00000048c744241003000000488b0c24488b542408eb0048b8abaaaaaaaaaaaaaa4889d348f7ea4801da48d1fa48c1fb3f4829da4801d148894c2428488b6c24184883c420c3",
		inputDir + "1_test_input.txt": "48c74424080000000048c74424082f010000c3",
		inputDir + "2_test_input.txt": "48c74424080000000048c744240800000000c3",
		inputDir + "3_test_input.txt": "0f57c00f11442410488b4424088400488b004889442408e90000000065488b0c2500000000488d842478feffff483b41100f863e0500004881ec080200004889ac2400020000488dac240002000048c7842410020000000000000f57c90f118c241802000048c7842408010000000000000f57c90f118c2410010000488b0500000000488d0d0000000048898c24e800000048898424f000000048898c24b800000048898424c000000048c744247000000000488b8424b8000000488b8c24c000000048898424a800000048898c24b000000048c74424500010000048c74424680000000048c784248800000000000000c644244600b48c78424a000000000000000488b8424b0000000488b8c24a8000000488d15000000004839d17405e949040000b901000000eb0048898424a0000000884c2447488b8424a000000048898424880000000fb6442447884424460fb64424468844244784c07505e906040000488b8424880000008400488b4008488944246048394424500f9e442447eb00807c2447007505e9e1010000488b8424880000004889442468eb00488b44246848898424980000004889442470eb004889842480000000feb00eb000f57c00f118424f8000000488b84248000000048890424e8000000008b4424088944244c488b442418488b4c242048898424f800000048898c24000100008b44244c89442448488b8424f8000000488b8c240001000048898424d800000048898c24e00000004883bc24d8000000000f95c0884424477505e930010000488b050000000048398424d80000000f94c07405e912010000488b8424d800000048890424488b8424e00000004889442408488b05000000004889442410e8000000000fb6442418eb0088442447eb00807c2447007502eb42beb00488b842408010000488b8c2410010000488b942418010000488984241002000048898c24180200004889942420020000488bac24000200004881c408020000c3eb00488b842408010000488b8c2410010000488b942418010000488d59014839d37602eb28eb008b742448893488488984240801000048899c24100100004889942418010000eb00e999feffff48894c2458488d350000000048893424488944240848894c2410488954241848895c2420e800000000488b442428488b4c2430488b542438488d5901488b4c2458eb97e91affffffe91bffffffeb0048837c2450107c05e9e601000048c744245010000000eb0048c74424780000000048c78424a801000000000000488dbc24b00100000f57c0488d7fd048896c24f0488d6c24f0e800000000488b6d00488d8424a801000048894424784889842490000000488d050000000048890424488b4424504889442408488b4424504889442410e800000000488b442418488b4c2450488b542450488984243801000048898c24400100004889942448010000488984242001000048898c24280100004889942430010000488b8424a8000000488b8c24b000000048898424c800000048898c24d000000048c784245001000000000000488dbc24580100000f57c0488d7fd048896c24f0488d6c24f0e800000000488b6d00488b842420010000488b8c2428010000488b942430010000488984245001000048898c24580100004889942460010000488b8424c8000000488b8c24d0000000488984246801000048898c247001000048c7842498010000ffffffff48c78424a0010000ffffffff488b8424900000008400833d00000000007402eb3d488b8c2450010000488908488d7808488db4245801000048896c24f0488d6c24f0e800000000488b6d00eb00eb00488b4424784889442468e95dfcffff488d0d0000000048890c244889442408488d8424500100004889442410e800000000ebcbe920feffffe914fcffff31c031c9be9b5fbffffe800000000e99dfaffff65488b0c2500000000488d8424f0feffff483b41100f86cc0300004881ec900100004889ac2488010000488dac24880100000f57c00f118424a801000048c7842460010000000000000f57c00f118424680100000f57c00f11842478010000488d050000000048898424a800000048c78424b00000000a00000048c7442450000000000f57c00f1184248800000048c7442468000000000f57c00f118424c800000048c7842480000000000000000f57c00f11842408010000488b8424a8000000488b8c24b00000004889042448894c240866c7442410d400c644241200e800000000488b4424184889842480000000488b442428488b4c242048898c24080100004889842410010000488b8424800000004889442468488b842410010000488b8c240801000048898c24c800000048898424d0000000488b4424684889442450488b8424d0000000488b8c24c800000048898c24880000004889842490000000eb00488b4424504889442478488944245848890424488b8424a0010000488b8c249801000048894c2408488944241048c7442418ffffffffe800000000488b442428488b4c2420488b54243048898c241801000048898424200100004889942428010000f48898c24480100004889842450010000488994245801000048c744244000000000488b8424500100004889442438488b8424480100004889442470488b44243848394424407c05e9cb010000eb00488b4424708400488b08488b400848898c24e800000048898424f000000048898c24b800000048898424c0000000b48890c244889442408488d0500000000488944241048c744241801000000e800000000488b442420488b4c242848898424f800000048898c2400010000488d94246001000048891424488944240848894c2410e800000000eb00f488b54244048ffc2488954244048395424387f02eb13488b4424704883c0104889442470e94fffffffeb00488d94246001000048895424600f57c00f1184249800000048837c2460007505e9d4000000488b5424608402488b42184889442448488b5424608402488b1a488b4a08488b5210eb004839c87605e9c80000004829c24889d648f7da48c1fa3f4821c24801da48899424300100004829c148898c24380100004889b4244001000048c7042400000000488954240848894c24104889742418e800000000488b442420488b4c242848898424d800000048898c24e0000000488984249800000048898c24a0000000eb00488b842498000000488b8c24a000000048898424a801000048898c24b0010000488bac24880100004881c490010000c3488d0500000000488984249800000048c78424a000000005000000ebb3fe9e5feffffe80000000090e800000000e90ffcffff"}

	for path, expected := range inputs {
		fmt.Println("case: ", path)
		file, _ := ioutil.ReadFile(path)
		parsed := sgo.Parse(string(file))
		assert.Equal(t, expected, parsed)
	}
}
