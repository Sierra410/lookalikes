package lookalikes

import "testing"

func TestDecode(t *testing.T) {
	got := Decode("Тне 🅠🅤🅘🅒🅚 𝐛𝐫𝐨𝐰𝐧 𝚏𝚘𝚡 ⒥⒰⒨⒫⒮ 🄾🅅🄴🅁 тнє 𝓵𝓪𝔃𝔂 𝖉𝖔𝖌 0𝟙ƻӟ４５𝟲𝟳𝟴𝟵")
	if got != "The quick brown fox jumps over the lazy dog 0123456789" {
		t.Errorf("Unexpected result: %s", got)
	}
}
