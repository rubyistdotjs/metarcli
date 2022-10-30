# Metar CLI

A Command Line Interface written in [Go](https://go.dev/) (with [Cobra](https://github.com/spf13/cobra)) that displays the latest METAR and TAF messages for one or multiple airports using the [CheckWX API](https://www.checkwxapi.com/).

## Installation

```
go install github.com/rubyistdotjs/metarcli@latest
```

## Usage

As of now there is only one command `metarcli icaoCodes... [flags]`.

### Flags
- `--apiKey` (`-k`) mandatory, your CheckWX API key.
- `--help` (`-h`)

### Examples

Lets say your are based in Montpellier, France (LFMT), you can retrieve the latest METAR and TAF with the following:
```bash
> metarcli LFMT --apiKey abcdefghijklmno0123456789

Montpellier-Méditerranée Airport
LFMT 301330Z AUTO 13006KT 9999 FEW015 BKN027 BKN032 21/19 Q1021 NOSIG
TAF LFMT 300800Z 3009/3109 04008KT 9999 SCT015 BKN020 BECMG 3010/3012 12008KT TEMPO 3022/3109 3000 BR OVC005

```

More useful, you can pass up-to 10 ICAO airport codes to get, for example, the weather of the surrounding airports or of the ones on your route:
```bash
> metarcli LFMT LFNG LFTW LFMV LFNH --apiKey abcdefghijklmno0123456789

Montpellier-Méditerranée Airport
LFMT 301330Z AUTO 13006KT 9999 FEW015 BKN027 BKN032 21/19 Q1021 NOSIG
TAF LFMT 300800Z 3009/3109 04008KT 9999 SCT015 BKN020 BECMG 3010/3012 12008KT TEMPO 3022/3109 3000 BR OVC005

Aérodrome de Montpellier - Candillargues
-
-

Nîmes-Arles-Camargue Airport
LFTW 301330Z AUTO 18008KT 160V230 9999 SCT025 23/16 Q1021 NOSIG
TAF LFTW 300800Z 3009/3109 VRB02KT 9999 BKN019

Avignon-Caumont Airport
LFMV 301330Z AUTO 25006KT 220V280 CAVOK 25/13 Q1021 NOSIG
TAF LFMV 300800Z 3009/3109 VRB02KT CAVOK TEMPO 3100/3106 0800 FG VV///

Aérodrome de Carpentras
-
-

```

LFNG (Candillargues) and LFNH (Carpentras) are small aerodrome without weather station, in this case the messages are replaced with a simple dash (`-`).

## License
[MIT](./LICENSE)
