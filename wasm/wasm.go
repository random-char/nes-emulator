//go:build wasm
// +build wasm

package main

import (
	"fmt"
	"nes-emulator/pkg/cartridge/loader/b64string"
	"nes-emulator/pkg/nes"
	"nes-emulator/pkg/video"
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "canvas")

	canvasVideoReceiver := video.NewCanvasVideoReceiver(canvas)

	nes := nes.New(canvasVideoReceiver)
	cartridge, err := b64string.LoadFromBase64(b64c)
	if err != nil {
		panic(err)
	}

	nes.InsertCartridge(cartridge)

	js.Global().Set("clock", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("clock")
		nes.Clock()
		return nil
	}))

	js.Global().Set("start", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("start")
		go nes.Start()
		return nil
	}))

	js.Global().Set("stop", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("stop")
		nes.Stop()
		return nil
	}))

	// js.Global().Set("frame", js.FuncOf(func(this js.Value, args []js.Value) any {
	// 	fmt.Println("frame")
	// 	nes.Frame()
	// 	return nil
	// }))

	for {
		select {}
	}
}

var b64c string = "TkVTGgEBAAAAAAAAAAAAAEz1xWB42KL/mq0CIBD7rQIgEPupAI0AII0BII0FII0FIK0CIKIgjgYgogCOBiCiAKAPqQCNByDK0PqI0PepP40GIKkAjQYgogC9eP+NByDo4CDQ9anAjRdAqQCNFUCpeIXQqfuF0al/hdOgAIwGIIwGIKkAhdepB4XQqcOF0SCnwiCNwqISIGHCpdVKSkqwHEqwDEqwJ0qwA0yBwEwmwSBvxsbXENupDYXX0NUgb8bm16XXyQ6QyqkAhdfwxCCJxqXX8AYg7cBMgcCpAIXY5tcg7cDm16XXyQ7Q9akAhdel2PACqf+FACDtwUyBwKXXCqq9CsGNAAK9C8GNAQKpwUip3kipAIUAbAACLcctx9vHhcjey/jN7s6iz3TR+9TUwUrfuNuq4akAhdepkoXQqcSF0SCnwiCNwqIPIGHCpdVKSkqwHEqwDEqwJ0qwA0w1wUxywCBvxsbXENupCoXX0NUgb8bm16XXyQuQyqkAhdfwxCCJxqXX8AYgocFMNcGpAIXY5tcgocHm16XXyQvQ9akAhdel2PACqf+FACDtwUw1waXXCqq9vsGNAAK9v8GNAQKpwUip3kipAIUAbAACo8ajxh7lPefT6Bbphuv27Wbw1vJG9akAhQAgANkg4Nrq6uqlAPAChdhM7cFMgcAgjcKpAIXTpdcYaQQKJtMKJtMKJtMKJtMKJtNIpdMJII0GIGgJBI0GIKUA8B3J//AmSkpKSqq9UcKNByClACkPqr1Rwo0HIEyUwqlPjQcgqUuNByBMlMKpRY0HIKlyjQcgTJTCMDEyMzQ1Njc4OUFCQ0RFRqXXGGkEqKmEjQAgqSCNBiCpAo0GIKkgiMjQAqkqjQcgiMrQ8amAjQAgTJTCpdLF0vD8YKkAjQUgjQUgqQCNBiCpAI0GIGCpAI0AII0BICDtwqkgjQYgoACMBiCiILHQ8CDJ//ANjQcgyNAC5tHK0O3w6cjQAubRqSCNByDK0Pjw2qmAjQAgqQ6NASBgqSCNBiCpAI0GIKIeqSCgII0HIIjQ+srQ9WD/////ICAgIC0tIFJ1biBhbGwgdGVzdHP/ICAgIC0tIEJyYW5jaCB0ZXN0c/8gICAgLS0gRmxhZyB0ZXN0c/8gICAgLS0gSW1tZWRpYXRlIHRlc3Rz/yAgICAtLSBJbXBsaWVkIHRlc3Rz/yAgICAtLSBTdGFjayB0ZXN0c/8gICAgLS0gQWNjdW11bGF0b3IgdGVzdHP/ICAgIC0tIChJbmRpcmVjdCxYKSB0ZXN0c/8gICAgLS0gWmVyb3BhZ2UgdGVzdHP/ICAgIC0tIEFic29sdXRlIHRlc3Rz/yAgICAtLSAoSW5kaXJlY3QpLFkgdGVzdHP/ICAgIC0tIEFic29sdXRlLFkgdGVzdHP/ICAgIC0tIFplcm9wYWdlLFggdGVzdHP/ICAgIC0tIEFic29sdXRlLFggdGVzdHP///8gICAgVXAvRG93bjogc2VsZWN0IHRlc3T/ICAgICAgU3RhcnQ6IHJ1biB0ZXN0/yAgICAgU2VsZWN0OiBJbnZhbGlkIG9wcyH/AP////8gICAgLS0gUnVuIGFsbCB0ZXN0c/8gICAgLS0gTk9QIHRlc3Rz/yAgICAtLSBMQVggdGVzdHP/ICAgIC0tIFNBWCB0ZXN0c/8gICAgLS0gU0JDIHRlc3QgKG9wY29kZSAwRUJoKf8gICAgLS0gRENQIHRlc3Rz/yAgICAtLSBJU0IgdGVzdHP/ICAgIC0tIFNMTyB0ZXN0c/8gICAgLS0gUkxBIHRlc3Rz/yAgICAtLSBTUkUgdGVzdHP/ICAgIC0tIFJSQSB0ZXN0c////////yAgICBVcC9Eb3duOiBzZWxlY3QgdGVzdP8gICAgICBTdGFydDogcnVuIHRlc3T/ICAgICBTZWxlY3Q6IE5vcm1hbCBvcHP/AEiKSK0CIKkgjQYgqUCNBiDm0qkAjQUgjQUgqQCNBiCpAI0GIKIJjhZAyo4WQK0WQEom1MrQ96XUqkXWJdSF1YbWaKpoQECiAIYAhhCGESAtxyDbxyCFyCDeyyD4zSDuziCizyB00SD71CAA2aUAhRCpAIUAIODaIErfILjbIKrhIKPGIB7lID3nINPoIBbpIIbrIPbtIGbwINbypQCFEakAhQAgRvWlAAUQBRHwDiBvxqYAhgKmEIYDTG7GIInGYKkDjRVAqYeNAECpiY0BQKnwjQJAqQCNA0BgqQKNFUCpP40EQKmajQVAqf+NBkCpAI0HQGCgTqn/hQEgsMYgt8Zgqf9IqarQBak0SKlVKASpRKlkqerq6uoISAypqerq6uoISBSpNKlUqXSp1Kn0qerq6uoISBo6Wnra+oCJ6urq6ghIHKmpPKmpXKmpfKmp3Kmp/Kmp6urq6ghIogVoyVXwCsmq8AZohABMKMdoKcvJAPAGycvwAoQAyMrQ4GDqOLAEogGGAOoYsANMQMeiAoYA6jiQA0xLx6IDhgDqGJAEogSGAOqpAPAEogWGAOqpQPADTGjHogaGAOqpQNAEogeGAOqpANADTH3HogiGAOqp/4UBJAFwBKIJhgDqJAFQA0yWx6IKhgDqqQCFASQBUASiC4YA6iQBcANMr8eiDIYA6qkAEASiDYYA6qmAEANM2ceiDoYA6qmAMASiD4YA6qkAMANM2ceiEIYA6mDqqf+FASQBqQA4ePgIaCnvyW/wBKIRhgDqqUCFASQB2KkQGAhoKe/JZPAEohKGAOqpgIUBJAH4qQA4CGgp78kv8ASiE4YA6qn/SCjQCRAHUAWQA0w1yKIUhgDqqQRIKPAJMAdwBbADTEnIohWGAOr4qf+FASQBGKkASKn/aNAJMAdQBbADTGfIohaGAOqpAIUBJAE4qf9IqQBo8AkQB3AFkANMhMiiF4YAYOoYqf+FASQBqVUJqrALEAnJ/9AFUANMosiiGIYA6ji4qQAJANAJcAeQBTADTLjIohmGAOoYJAGpVSmq0AlQB7AFMANMz8iiGoYA6ji4qfgp75ALEAnJ6NAFcANM58iiG4YA6hgkAalfSaqwCxAJyfXQBVADTADJohyGAOo4uKlwSXDQCXAHkAUwA0wWyaIdhgDqGCQBqQBpaTALsAnJadAFcANML8miHoYA6jj4JAGpAWlpMAuwCclr0AVwA0xJyaIfhgDq2Di4qX9pfxALsAnJ/9AFUANMYsmiIIYA6hgkAal/aYAQC7AJyf/QBXADTHvJoiGGAOo4uKl/aYDQCTAHcAWQA0yRyaIihgDqOLipn/AJEAdwBZADTKXJoiOGAOoYJAGpANAJMAdQBbADTLrJoiOGAOokAalAyUAwCZAH0AVQA0zQyaIkhgDquMk/8AkwB5AFcANM48miJYYA6slB8AcQBRADTPPJoiaGAOqpgMkA8AcQBZADTAXKoieGAOrJgNAHMAWQA0wVyqIohgDqyYGwB/AFEANMJcqiKYYA6sl/kAfwBTADTDXKoiqGAOokAaBAwEDQCTAHkAVQA0xLyqIrhgDquMA/8AkwB5AFcANMXsqiLIYA6sBB8AcQBRADTG7Koi2GAOqggMAA8AcQBZADTIDKoi6GAOrAgNAHMAWQA0yQyqIvhgDqwIGwB/AFEANMoMqiMIYA6sB/kAfwBTADTLDKojGGAOokAaJA4EDQCTAHkAVQA0zGyqkyhQDquOA/8AkwB5AFcANM2cqpM4UA6uBB8AcQBRADTOnKqTSFAOqigOAA8AcQBZADTPvKqTWFAOrggNAHMAWQA0wLy6k2hQDq4IGwB/AFEANMG8upN4UA6uB/kAfwBTADTCvLqTiFAOo4uKKf8AkQB3AFkANMP8uiOYYA6hgkAaIA0AkwB1AFsANMVMuiOoYA6ji4oJ/wCRAHcAWQA0xoy6I7hgDqGCQBoADQCTAHUAWwA0x9y6I8hgDqqVWiqqAzyVXQI+Cq0B/AM9AbyVXQF+Cq0BPAM9APyVbwC+Cr8AfANPADTK/Loj2GAKBxIDH56UAgN/nIIEf56T8gTPnIIFz56UEgYvnIIHL56QAgdvnIIID56X8ghPlg6qn/hQGpRKJVoGboiOBW0CHAZdAd6OiIiOBY0BXAY9ARysjgV9ALwGTQB8lE0ANMFMyiPoYA6jiiaamWJAGg/8jQPTA7kDlQN8AA0DPI8DAwLpAsUCoYuKAAiPAjECGwH3AdwP/QGRiI8BUQE7ARcA/A/tALyZbQB+Bp0ANMYsyiP4YA6jigaamWJAGi/+jQPTA7kDlQN+AA0DPo8DAwLpAsUCoYuKIAyvAjECGwH3Ad4P/QGRjK8BUQE7ARcA/g/tALyZbQB8Bp0ANMsMyiQIYA6qmFojSgmRgkAajwLrAsUCoQKMmF0CTgNNAgwIXQHKkAOLio0BWQE3ARMA/JANAL4DTQB8AA0ANM78yiQYYA6qmFojSgmRgkAarwLrAsUCoQKMmF0CTghdAgwJnQHKkAOLiq0BWQE3ARMA/JANAL4ADQB8CZ0ANMLs2iQoYA6qmFojSgmRgkAZjwLrAsUCoQKMmZ0CTgNNAgwJnQHKAAOLiY0BWQE3ARMA/JANAL4DTQB8AA0ANMbc2iQ4YA6qmFojSgmRgkAYrwLrAsUCowKMk00CTgNNAgwJnQHKIAOLiK0BWQE3ARMA/JANAL4ADQB8CZ0ANMrM2iRIYA6rqO/wegM6JpqYQYJAGa8DIQMLAuUCzJhNAo4GnQJMAz0CCgAakEOLiiALrwFTATkBFwD+Bp0AvJBNAHwAHQA0zzzaJFhgCu/weaYKn/hQG6jv8H6qKAmqkzSKlpSLrgftAgaMlp0BtoyTPQFrrggNARrYAByTPQCq1/Aclp0ANMM86iRoYA6qKAmiA9zkxbzrrgftAZaGi64IDQEqkAIE7OaMlN0Ahoyc7QA0xfzqJHhgDqqc5IqWZIYKJ3oGkYJAGpgyBmzvAkECKwIFAeyYPQGsBp0Bbgd9ASOLipACBmztAJMAeQBXADTJ3OokiGAOqpzkiprkipZUipVaCIoplAMDVQM/AxkC/JVdArwIjQJ+CZ0COpzkipzkiph0ipVUAQFXAT0BGQD8lV0AvAiNAH4JnQA0zpzqJJhgCu/weaYKJVoGmp/4UB6iQBOKkBSpAd0BswGVAXyQDQE7ipqkqwDfALMAlwB8lV0ANMIM+iSoYA6iQBOKmACpAe0BwwGlAYyQDQFLg4qVUKsA3wCxAJcAfJqtADTEvPokuGAOokATipAWqQHvAcEBpQGMmA0BS4GKlVapAN8AswCXAHySrQA0x2z6JMhgDqJAE4qYAqkB7wHDAaUBjJAdAUuBipVSqwDfALEAlwB8mq0ANMoc+iTYYAYKUAjf8HqQCFgKkChYGp/4UBqQCFgqkDhYOFhKkAhf+pBIUAqVqNAAKpW40AA6lcjQMDqV2NAASiAKGAyVrQH+jooYDJW9AX6KGAyVzQEKIAof/JXdAIooGh/8la8AWpWI3/B6mqogCBgOjoqauBgOiprIGAogCprYH/rQACyarQFa0AA8mr0A6tAwPJrNAHrQAEya3wBalZjf8Hrf8HhQCpAI0AA6mqjQACogCgWiC29wGAIMD3yCDO9wGCINP3yCDf9yGAIOX3yKnvjQADIPH3IYIg9vfIIAT4QYAgCvjIqXCNAAMgGPhBgiAd+MipaY0AAiAp+GGAIC/4yCA9+GGAIEP4yKl/jQACIFH4YYAgVvjIqYCNAAIgZPhhgCBq+MggePhhgCB9+MipQI0AAiCJ+MGAII74yEipP40AAmggmvjBgCCc+MhIqUGNAAJowYAgqPjISKkAjQACaCCy+MGAILX4yEipgI0AAmjBgCC/+MhIqYGNAAJowYAgyfjISKl/jQACaMGAINP4yKlAjQACIDH54YAgN/nIqT+NAAIgR/nhgCBM+cipQY0AAiBc+eGAIGL5yKkAjQACIHL54YAgdvnIqX+NAAIggPnhgCCE+WCpVYV4qf+FASQBoBGiI6kApXjwEDAOyVXQCsAR0AbgI1AC8ASpdoUAqUYkAYV48AoQCFAGpXjJRvAEqXeFAKlVhXgkAakRoiOgAKR48BAwDsBV0ArJEdAG4CNQAvAEqXiFAKBGJAGEePAKEAhQBqR4wEbwBKl5hQAkAalVhXigEakjogCmePAQMA7gVdAKwBHQBskjUALwBKl6hQCiRiQBhnjwChAIUAameOBG8ASpe4UAqcCFeKIzoIipBSR4EBBQDtAMyQXQCOAz0ATAiPAEqXyFAKkDhXipASR4MAhwBvAEyQHwBKl9hQCgfqmqhXggtvcFeCDA98ipAIV4IM73BXgg0/fIqaqFeCDf9yV4IOX3yKnvhXgg8fcleCD298ipqoV4IAT4RXggCvjIqXCFeCAY+EV4IB34yKlphXggKfhleCAv+MggPfhleCBD+Mipf4V4IFH4ZXggVvjIqYCFeCBk+GV4IGr4yCB4+GV4IH34yKlAhXggifjFeCCO+MhIqT+FeGggmvjFeCCc+MhIqUGFeGjFeCCo+MhIqQCFeGggsvjFeCC1+MhIqYCFeGjFeCC/+MhIqYGFeGjFeCDJ+MhIqX+FeGjFeCDT+MipQIV4IDH55XggN/nIqT+FeCBH+eV4IEz5yKlBhXggXPnleCBi+cipAIV4IHL55XggdvnIqX+FeCCA+eV4IIT5yKlAhXggifiq5HggjvjIqT+FeCCa+OR4IJz4yKlBhXjkeCCo+MipAIV4ILL4quR4ILX4yKmAhXjkeCC/+MipgYV45HggyfjIqX+FeOR4INP4yJiqqUCFeCDd+MR4IOL46Kk/hXgg7vjEeCDw+OipQYV4xHgg/PjoqQCFeCAG+cR4IAn56KmAhXjEeCAT+eipgYV4xHggHfnoqX+FeMR4ICf56IqoIJD5hXhGeKV4IJ35yIV4RnileCCt+cggvfmFeAZ4pXggw/nIhXgGeKV4INT5yCDk+YV4ZnileCDq+ciFeGZ4pXgg+/nIIAr6hXgmeKV4IBD6yIV4JnileCAh+qn/hXiFASQBOOZ40AwwClAIkAaleMkA8ASpq4UAqX+FeLgY5njwDBAKcAiwBqV4yYDwBKmshQCpAIV4JAE4xnjwDBAKUAiQBqV4yf/wBKmthQCpgIV4uBjGePAMMApwCLAGpXjJf/AEqa6FAKkBhXjGePAEqa+FAGCpVY14Bqn/hQEkAaARoiOpAK14BvAQMA7JVdAKwBHQBuAjUALwBKmwhQCpRiQBjXgG8AsQCVAHrXgGyUbwBKmxhQCpVY14BiQBqRGiI6AArHgG8BAwDsBV0ArJEdAG4CNQAvAEqbKFAKBGJAGMeAbwCxAJUAeseAbARvAEqbOFACQBqVWNeAagEakjogCueAbwEDAO4FXQCsAR0AbJI1AC8ASptIUAokYkAY54BvALEAlQB654BuBG8ASptYUAqcCNeAaiM6CIqQUseAYQEFAO0AzJBdAI4DPQBMCI8ASptoUAqQONeAapASx4BjAIcAbwBMkB8ASpt4UAoLipqo14BiC29w14BiDA98ipAI14BiDO9w14BiDT98ipqo14BiDf9y14BiDl98ip7414BiDx9y14BiD298ipqo14BiAE+E14BiAK+MipcI14BiAY+E14BiAd+MipaY14BiAp+G14BiAv+MggPfhteAYgQ/jIqX+NeAYgUfhteAYgVvjIqYCNeAYgZPhteAYgavjIIHj4bXgGIH34yKlAjXgGIIn4zXgGII74yEipP414BmggmvjNeAYgnPjISKlBjXgGaM14BiCo+MhIqQCNeAZoILL4zXgGILX4yEipgI14BmjNeAYgv/jISKmBjXgGaM14BiDJ+MhIqX+NeAZozXgGINP4yKlAjXgGIDH57XgGIDf5yKk/jXgGIEf57XgGIEz5yKlBjXgGIFz57XgGIGL5yKkAjXgGIHL57XgGIHb5yKl/jXgGIID57XgGIIT5yKlAjXgGIIn4qux4BiCO+MipP414BiCa+Ox4BiCc+MipQY14Bux4BiCo+MipAI14BiCy+KrseAYgtfjIqYCNeAbseAYgv/jIqYGNeAbseAYgyfjIqX+NeAbseAYg0/jImKqpQI14BiDd+Mx4BiDi+OipP414BiDu+Mx4BiDw+OipQY14Bsx4BiD8+OipAI14BiAG+cx4BiAJ+eipgI14Bsx4BiAT+eipgY14Bsx4BiAd+eipf414Bsx4BiAn+eiKqCCQ+Y14Bk54Bq14BiCd+ciNeAZOeAateAYgrfnIIL35jXgGDngGrXgGIMP5yI14Bg54Bq14BiDU+cgg5PmNeAZueAateAYg6vnIjXgGbngGrXgGIPv5yCAK+o14Bi54Bq14BiAQ+siNeAYueAateAYgIfqp/414BoUBJAE47ngG0A0wC1AJkAeteAbJAPAEqeWFAKl/jXgGuBjueAbwDRALcAmwB614BsmA8ASp5oUAqQCNeAYkATjOeAbwDRALUAmQB614Bsn/8ASp54UAqYCNeAa4GM54BvANMAtwCbAHrXgGyX/wBKnohQCpAY14Bs54BvAEqemFAGCpo4UzqYmNAAOpEo1FAqn/hQGiZakAhYmpA4WKoAA4qQC4sYnwDJAKcAjJidAE4GXwBKnqhQCp/4WXhZgkmKA0sZfJo9ACsASp64UApQBIqUaF/6kBhQCg/7H/yRLwBGip7EhohQCi7akAhTOpBIU0oAAYqf+FASQBqaqNAASpVREzsAgQBsn/0AJwAoYA6Di4qQARM/AGcASQAjAChgDoGCQBqVUxM9AGUASwAhAChgDoOLip740ABKn4MTOQCBAGyejQAlAChgDoGCQBqaqNAASpX1EzsAgQBsn10AJwAoYA6Di4qXCNAARRM9AGcASQAhAChgDoGCQBqWmNAASpAHEzMAiwBslp0AJQAoYA6DgkAakAcTMwCLAGyWrQAlAChgDoOLipf40ABHEzEAiwBsn/0AJwAoYA6BgkAamAjQAEqX9xMxAIsAbJ/9ACUAKGAOg4uKmAjQAEqX9xM9AGMARwArAChgDoJAGpQI0ABNEzMAaQBNACcAKGAOi4zgAE0TPwBjAEkAJQAoYA6O4ABO4ABNEz8AIwAoYA6KkAjQAEqYDRM/AEEAKwAoYA6KCAjAAEoADRM9AEMAKwAoYA6O4ABNEzsATwAjAChgDozgAEzgAE0TOQBPACEAKGAGCpAIUzqQSFNKAAogEkAalAjQAEOPEzMAqQCNAGcATJAPAChgDouDipQM4ABPEz8AowCJAGcATJAfAChgDoqUA4JAHuAATuAATxM7AK8AgQBnAEyf/wAoYA6BipAI0ABKmA8TOQBMl/8AKGAOg4qX+NAASpgfEzUAaQBMkC8AKGAOipAKmHkTOtAATJh/AChgDoqX6NAAKp240BAmwAAqkAjf8CqQGNAAOpA40AAqmpjQABqVWNAQGpYI0CAampjQADqaqNAQOpYI0CAyC128mq8AKGAGBs/wKp/4UBqaqFM6m7hYmiAKlmJAE4oAC0MxAS8BBQDpAMyWbQCOAA0ATAqvAEqQiFAKKKqWa4GKAAtP8QEvAQcA6wDMC70AjJZtAE4IrwBKkJhQAkATigRKIAlDOlM5AYyUTQFFASGLigmaKAlIWlBbAGyZnQAlAEqQqFAKALqaqieIV4ILb3FQAgwPfIqQCFeCDO9xUAINP3yKmqhXgg3/c1ACDl98ip74V4IPH3NQAg9vfIqaqFeCAE+FUAIAr4yKlwhXggGPhVACAd+MipaYV4ICn4dQAgL/jIID34dQAgQ/jIqX+FeCBR+HUAIFb4yKmAhXggZPh1ACBq+MggePh1ACB9+MipQIV4IIn41QAgjvjISKk/hXhoIJr41QAgnPjISKlBhXho1QAgqPjISKkAhXhoILL41QAgtfjISKmAhXho1QAgv/jISKmBhXho1QAgyfjISKl/hXho1QAg0/jIqUCFeCAx+fUAIDf5yKk/hXggR/n1ACBM+cipQYV4IFz59QAgYvnIqQCFeCBy+fUAIHb5yKl/hXgggPn1ACCE+amqhTOpu4WJogCgZiQBOKkAtTMQEvAQUA6QDMBm0AjgANAEyarwBKkihQCiiqBmuBipALX/EBLwEHAOsAzJu9AIwGbQBOCK8ASpI4UAJAE4qUSiAJUzpTOQGMlE0BRQEhi4qZmigJWFpQWwBsmZ0AJQBKkkhQCgJaJ4IJD5lQBWALUAIJ35yJUAVgC1ACCt+cggvfmVABYAtQAgw/nIlQAWALUAINT5yCDk+ZUAdgC1ACDq+ciVAHYAtQAg+/nIIAr6lQA2ALUAIBD6yJUANgC1ACAh+qn/lQCFASQBOPYA0AwwClAIkAa1AMkA8ASpLYUAqX+VALgY9gDwDBAKcAiwBrUAyYDwBKkuhQCpAJUAJAE41gDwDBAKUAiQBrUAyf/wBKkvhQCpgJUAuBjWAPAMMApwCLAGtQDJf/AEqTCFAKkBlQDWAPAEqTGFAKkzhXipRKB4ogA4JAG2AJASUBAwDvAM4DPQCMB40ATJRPAEqTKFAKmXhX+pR6D/ogAYuLaAsBJwEBAO8Azgl9AIwP/QBMlH8ASpM4UAqQCFf6lHoP+iaRi4loCwGHAWMBTwEuBp0A7A/9AKyUfQBqV/yWnwBKk0hQCp9YVPqUegTyQBogA4lgCQFlAUMBLQEOAA0AzAT9AIyUfQBKVP8ASpNYUAYKmJjQADqaOFM6kSjUUComWgADipALi5AAPwDJAKcAjJidAE4GXwBKk2hQCp/4UBJAGgNLn//8mj0AKwBKk3hQCpRoX/oP+5RgHJEvAEqTiFAKI5GKn/hQEkAamqjQAEqVWgABkABLAIEAbJ/9ACcAKGAOg4uKkAGQAE8AZwBJACMAKGAOgYJAGpVTkABNAGUASwAhAChgDoOLip740ABKn4OQAEkAgQBsno0AJQAoYA6BgkAamqjQAEqV9ZAASwCBAGyfXQAnAChgDoOLipcI0ABFkABNAGcASQAhAChgDoGCQBqWmNAASpAHkABDAIsAbJadACUAKGAOg4JAGpAHkABDAIsAbJatACUAKGAOg4uKl/jQAEeQAEEAiwBsn/0AJwAoYA6BgkAamAjQAEqX95AAQQCLAGyf/QAlAChgDoOLipgI0ABKl/eQAE0AYwBHACsAKGAOgkAalAjQAE2QAEMAaQBNACcAKGAOi4zgAE2QAE8AYwBJACUAKGAOjuAATuAATZAATwAjAChgDoqQCNAASpgNkABPAEEAKwAoYA6KCAjAAEoADZAATQBDACsAKGAOjuAATZAASwBPACMAKGAOjOAATOAATZAASQBPACEAKGAOgkAalAjQAEOPkABDAKkAjQBnAEyQDwAoYA6Lg4qUDOAAT5AATwCjAIkAZwBMkB8AKGAOipQDgkAe4ABO4ABPkABLAK8AgQBnAEyf/wAoYA6BipAI0ABKmA+QAEkATJf/AChgDoOKl/jQAEqYH5AARQBpAEyQLwAoYA6KkAqYeZAAStAATJh/AChgBgqf+FAamqjTMGqbuNiQaiAKlmJAE4oAC8MwYQEvAQUA6QDMlm0AjgANAEwKrwBKlRhQCiiqlmuBigALz/BRAS8BBwDrAMwLvQCMlm0ATgivAEqVKFAKBTqaqieI14BiC29x0ABiDA98ipAI14BiDO9x0ABiDT98ipqo14BiDf9z0ABiDl98ip7414BiDx9z0ABiD298ipqo14BiAE+F0ABiAK+MipcI14BiAY+F0ABiAd+MipaY14BiAp+H0ABiAv+MggPfh9AAYgQ/jIqX+NeAYgUfh9AAYgVvjIqYCNeAYgZPh9AAYgavjIIHj4fQAGIH34yKlAjXgGIIn43QAGII74yEipP414BmggmvjdAAYgnPjISKlBjXgGaN0ABiCo+MhIqQCNeAZoILL43QAGILX4yEipgI14BmjdAAYgv/jISKmBjXgGaN0ABiDJ+MhIqX+NeAZo3QAGINP4yKlAjXgGIDH5/QAGIDf5yKk/jXgGIEf5/QAGIEz5yKlBjXgGIFz5/QAGIGL5yKkAjXgGIHL5/QAGIHb5yKl/jXgGIID5/QAGIIT5qaqNMwapu42JBqIAoGYkATipAL0zBhAS8BBQDpAMwGbQCOAA0ATJqvAEqWqFAKKKoGa4GKkAvf8FEBLwEHAOsAzJu9AIwGbQBOCK8ASpa4UAJAE4qUSiAJ0zBq0zBpAayUTQFlAUGLipmaKAnYUFrQUGsAbJmdACUASpbIUAoG2ibSCQ+Z0ABl4ABr0ABiCd+cidAAZeAAa9AAYgrfnIIL35nQAGHgAGvQAGIMP5yJ0ABh4ABr0ABiDU+cgg5PmdAAZ+AAa9AAYg6vnInQAGfgAGvQAGIPv5yCAK+p0ABj4ABr0ABiAQ+sidAAY+AAa9AAYgIfqp/50ABoUBJAE4/gAG0A0wC1AJkAe9AAbJAPAEqXWFAKl/nQAGuBj+AAbwDRALcAmwB70ABsmA8ASpdoUAqQCdAAYkATjeAAbwDRALUAmQB70ABsn/8ASpd4UAqYCdAAa4GN4ABvANMAtwCbAHvQAGyX/wBKl4hQCpAZ0ABt4ABvAEqXmFAKkzjXgGqUSgeKIAOCQBvgAGkBJQEDAO8AzgM9AIwHjQBMlE8ASpeoUAqZeNfwapR6D/ogAYuL6ABbAScBAQDvAM4JfQCMD/0ATJR/AEqXuFAGCpVY2ABamqjTIEqYCFQ6kFhUSpMoVFqQSFRqIDoHep/4UBJAE4qQCjQOrq6urwEjAQUA6QDMlV0AjgVdAEwHfwBKl8hQCiBaAzuBipAKNA6urq6vASEBBwDrAMyarQCOCq0ATAM/AEqX2FAKmHhWepMoVooFckATipAKdn6urq6vASEBBQDpAMyYfQCOCH0ATAV/AEqX6FAKBTuBipAKdo6urq6vASMBBwDrAMyTLQCOAy0ATAU/AEqX+FAKmHjXcFqTKNeAWgVyQBOKkAr3cF6urq6vASEBBQDpAMyYfQCOCH0ATAV/AEqYCFAKBTuBipAK94Berq6urwEjAQcA6wDMky0AjgMtAEwFPwBKmBhQCp/4VDqQSFRKkyhUWpBIVGqVWNgAWpqo0yBKIDoIEkATipALND6urq6vASMBBQDpAMyVXQCOBV0ATAgfAEqYKFAKIFoAC4GKkAs0Xq6urq8BIQEHAOsAzJqtAI4KrQBMAA8ASpg4UAqYeFZ6kyhWigVyQBOKkAtxDq6urq8BIQEFAOkAzJh9AI4IfQBMBX8ASphIUAoP+4GKkAt2nq6urq8BIwEHAOsAzJMtAI4DLQBMD/8ASphYUAqYeNhwWpMo2IBaAwJAE4qQC/VwXq6urq8BIQEFAOkAzJh9AI4IfQBMAw8ASphoUAoEC4GKkAv0gF6urq6vASMBBwDrAMyTLQCOAy0ATAQPAEqYeFAGCpwIUBqQCNiQSpiYVgqQSFYaBEohepPiQBGINJ6urq6tAZsBdQFRATyT7QD8BE0AvgF9AHrYkEyRbwBKmIhQCgRKJ6qWY4uIPm6urq6vAZkBdwFTATyWbQD8BE0AvgetAHrYkEyWLwBKmJhQCp/4VJoESiqqlVJAEYh0nq6urq8BiwFlAUEBLJVdAOwETQCuCq0AalSckA8ASpioUAqQCFVqBYou+pZji4h1bq6urq8BiQFnAUMBLJZtAOwFjQCuDv0AalVslm8ASpi4UAqf+NSQWg5aKvqfUkARiPSQXq6urq8BmwF1AVEBPJ9dAPwOXQC+Cv0AetSQXJpfAEqYyFAKkAjVYFoFiis6mXOLiPVgXq6urq8BmQF3AVEBPJl9APwFjQC+Cz0AetVgXJk/AEqY2FAKn/hUmg/6KqqVUkARiXSurq6urwGLAWUBQQEslV0A7A/9AK4KrQBqVJyQDwBKmOhQCpAIVWoAai76lmOLiXUOrq6urwGJAWcBQwEslm0A7ABtAK4O/QBqVWyWbwBKmPhQBgoJAgMfnrQOrq6uogN/nIIEf56z/q6urqIEz5yCBc+etB6urq6iBi+cggcvnrAOrq6uogdvnIIID563/q6urqIIT5YKn/hQGglaICqUeFR6kGhUip641HBiAx+sNF6urq6iA3+q1HBsnq8AKEAMipAI1HBiBC+sNF6urq6iBH+q1HBsn/8AKEAMipN41HBiBU+sNF6urq6iBZ+q1HBsk28AKEAMip64VHIDH6x0fq6urqIDf6pUfJ6vAChADIqQCFRyBC+sdH6urq6iBH+qVHyf/wAoQAyKk3hUcgVPrHR+rq6uogWfqlR8k28AKEAMip641HBiAx+s9HBurq6uogN/qtRwbJ6vAChADIqQCNRwYgQvrPRwbq6urqIEf6rUcGyf/wAoQAyKk3jUcGIFT6z0cG6urq6iBZ+q1HBsk28AKEAKnrjUcGqUiFRakFhUag/yAx+tNF6uoISKCeaCggN/qtRwbJ6vAChACg/6kAjUcGIEL600Xq6ghIoJ9oKCBH+q1HBsn/8AKEAKD/qTeNRwYgVPrTRerqCEigoGgoIFn6rUcGyTbwAoQAoKGi/6nrhUcgMfrXSOrq6uogN/qlR8nq8AKEAMipAIVHIEL610jq6urqIEf6pUfJ//AChADIqTeFRyBU+tdI6urq6iBZ+qVHyTbwAoQAqeuNRwag/yAx+ttIBerqCEigpGgoIDf6rUcGyerwAoQAoP+pAI1HBiBC+ttIBerqCEigpWgoIEf6rUcGyf/wAoQAoP+pN41HBiBU+ttIBerqCEigpmgoIFn6rUcGyTbwAoQAoKei/6nrjUcGIDH630gF6urq6iA3+q1HBsnq8AKEAMipAI1HBiBC+t9IBerq6uogR/qtRwbJ//AChADIqTeNRwYgVPrfSAXq6urqIFn6rUcGyTbwAoQAYKn/hQGgqqICqUeFR6kGhUip641HBiCx+uNF6urq6iC3+q1HBsns8AKEAMip/41HBiDC+uNF6urq6iDH+q1HBskA8AKEAMipN41HBiDU+uNF6urq6iDa+q1HBsk48AKEAMip64VHILH650fq6urqILf6pUfJ7PAChADIqf+FRyDC+udH6urq6iDH+qVHyQDwAoQAyKk3hUcg1PrnR+rq6uog2vqlR8k48AKEAMip641HBiCx+u9HBurq6uogt/qtRwbJ7PAChADIqf+NRwYgwvrvRwbq6urqIMf6rUcGyQDwAoQAyKk3jUcGINT670cG6urq6iDa+q1HBsk48AKEAKnrjUcGqUiFRakFhUag/yCx+vNF6uoISKCzaCggt/qtRwbJ7PAChACg/6n/jUcGIML680Xq6ghIoLRoKCDH+q1HBskA8AKEAKD/qTeNRwYg1PrzRerqCEigtWgoINr6rUcGyTjwAoQAoLai/6nrhUcgsfr3SOrq6uogt/qlR8ns8AKEAMip/4VHIML690jq6urqIMf6pUfJAPAChADIqTeFRyDU+vdI6urq6iDa+qVHyTjwAoQAqeuNRwag/yCx+vtIBerqCEiguWgoILf6rUcGyezwAoQAoP+p/41HBiDC+vtIBerqCEigumgoIMf6rUcGyQDwAoQAoP+pN41HBiDU+vtIBerqCEigu2goINr6rUcGyTjwAoQAoLyi/6nrjUcGILH6/0gF6urq6iC3+q1HBsns8AKEAMip/41HBiDC+v9IBerq6uogx/qtRwbJAPAChADIqTeNRwYg1Pr/SAXq6urqINr6rUcGyTjwAoQAYKn/hQGgv6ICqUeFR6kGhUippY1HBiB7+gNF6urq6iCB+q1HBslK8AKEAMipKY1HBiCM+gNF6urq6iCR+q1HBslS8AKEAMipN41HBiCe+gNF6urq6iCk+q1HBslu8AKEAMippYVHIHv6B0fq6urqIIH6pUfJSvAChADIqSmFRyCM+gdH6urq6iCR+qVHyVLwAoQAyKk3hUcgnvoHR+rq6uogpPqlR8lu8AKEAMippY1HBiB7+g9HBurq6uoggfqtRwbJSvAChADIqSmNRwYgjPoPRwbq6urqIJH6rUcGyVLwAoQAyKk3jUcGIJ76D0cG6urq6iCk+q1HBslu8AKEAKmljUcGqUiFRakFhUag/yB7+hNF6uoISKDIaCgggfqtRwbJSvAChACg/6kpjUcGIIz6E0Xq6ghIoMloKCCR+q1HBslS8AKEAKD/qTeNRwYgnvoTRerqCEigymgoIKT6rUcGyW7wAoQAoMui/6mlhUcge/oXSOrq6uoggfqlR8lK8AKEAMipKYVHIIz6F0jq6urqIJH6pUfJUvAChADIqTeFRyCe+hdI6urq6iCk+qVHyW7wAoQAqaWNRwag/yB7+htIBerqCEigzmgoIIH6rUcGyUrwAoQAoP+pKY1HBiCM+htIBerqCEigz2goIJH6rUcGyVLwAoQAoP+pN41HBiCe+htIBerqCEig0GgoIKT6rUcGyW7wAoQAoNGi/6mljUcGIHv6H0gF6urq6iCB+q1HBslK8AKEAMipKY1HBiCM+h9IBerq6uogkfqtRwbJUvAChADIqTeNRwYgnvofSAXq6urqIKT6rUcGyW7wAoQAYKn/hQGg1KICqUeFR6kGhUippY1HBiBT+yNF6urq6iBZ+61HBslK8AKEAMipKY1HBiBk+yNF6urq6iBp+61HBslS8AKEAMipN41HBiBo+iNF6urq6iBu+q1HBslv8AKEAMippYVHIFP7J0fq6urqIFn7pUfJSvAChADIqSmFRyBk+ydH6urq6iBp+6VHyVLwAoQAyKk3hUcgaPonR+rq6uogbvqlR8lv8AKEAMippY1HBiBT+y9HBurq6uogWfutRwbJSvAChADIqSmNRwYgZPsvRwbq6urqIGn7rUcGyVLwAoQAyKk3jUcGIGj6L0cG6urq6iBu+q1HBslv8AKEAKmljUcGqUiFRakFhUag/yBT+zNF6uoISKDdaCggWfutRwbJSvAChACg/6kpjUcGIGT7M0Xq6ghIoN5oKCBp+61HBslS8AKEAKD/qTeNRwYgaPozRerqCEig32goIG76rUcGyW/wAoQAoOCi/6mlhUcgU/s3SOrq6uogWfulR8lK8AKEAMipKYVHIGT7N0jq6urqIGn7pUfJUvAChADIqTeFRyBo+jdI6urq6iBu+qVHyW/wAoQAqaWNRwag/yBT+ztIBerqCEig42goIFn7rUcGyUrwAoQAoP+pKY1HBiBk+ztIBerqCEig5GgoIGn7rUcGyVLwAoQAoP+pN41HBiBo+jtIBerqCEig5WgoIG76rUcGyW/wAoQAoOai/6mljUcGIFP7P0gF6urq6iBZ+61HBslK8AKEAMipKY1HBiBk+z9IBerq6uogafutRwbJUvAChADIqTeNRwYgaPo/SAXq6urqIG76rUcGyW/wAoQAYKn/hQGg6aICqUeFR6kGhUippY1HBiAd+0NF6urq6iAj+61HBslS8AKEAMipKY1HBiAu+0NF6urq6iAz+61HBskU8AKEAMipN41HBiBA+0NF6urq6iBG+61HBskb8AKEAMippYVHIB37R0fq6urqICP7pUfJUvAChADIqSmFRyAu+0dH6urq6iAz+6VHyRTwAoQAyKk3hUcgQPtHR+rq6uogRvulR8kb8AKEAMippY1HBiAd+09HBurq6uogI/utRwbJUvAChADIqSmNRwYgLvtPRwbq6urqIDP7rUcGyRTwAoQAyKk3jUcGIED7T0cG6urq6iBG+61HBskb8AKEAKmljUcGqUiFRakFhUag/yAd+1NF6uoISKDyaCggI/utRwbJUvAChACg/6kpjUcGIC77U0Xq6ghIoPNoKCAz+61HBskU8AKEAKD/qTeNRwYgQPtTRerqCEig9GgoIEb7rUcGyRvwAoQAoPWi/6mlhUcgHftXSOrq6uogI/ulR8lS8AKEAMipKYVHIC77V0jq6urqIDP7pUfJFPAChADIqTeFRyBA+1dI6urq6iBG+6VHyRvwAoQAqaWNRwag/yAd+1tIBerqCEig+GgoICP7rUcGyVLwAoQA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AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="