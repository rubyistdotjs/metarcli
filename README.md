# Metar CLI

A Command Line Interface written in [Go](https://go.dev/) (with [Cobra](https://github.com/spf13/cobra)) that displays the latest METAR and TAF messages for one or multiple airports using the [CheckWX API](https://www.checkwxapi.com/).

## Usage

As of now there is only one command `metar [options]` (see below for the list of options or you may use `metar --help`).

Lets say your are based in Montpellier, France (LFMT), you can retrieve the latest METAR and TAF with the following:
```bash
> metar --apiKey abcdefghijklmno0123456789 --icaoCodes LFMT

Montpellier-Méditerranée Airport
LFMT 070930Z AUTO 07005KT CAVOK 20/16 Q1025 NOSIG
TAF LFMT 070800Z 0709/0809 03008KT 9999 FEW015 TEMPO 0721/0809 -SHRA FEW065TCU BKN070

```

More useful, `icaoCodes` can take a list to get the weather of the surrounding airports or of the ones on your route. For example we plan to go from LFMT to LFNH:
```bash
> metar --apiKey abcdefghijklmno0123456789 --icaoCodes LFMT,LFNG,LFTW,LFMV,LFNH

Montpellier-Méditerranée Airport
LFMT 070930Z AUTO 07005KT CAVOK 20/16 Q1025 NOSIG
TAF LFMT 070800Z 0709/0809 03008KT 9999 FEW015 TEMPO 0721/0809 -SHRA FEW065TCU BKN070

Aérodrome de Montpellier - Candillargues
-
-

Nîmes-Arles-Camargue Airport
LFTW 070930Z AUTO 06005KT 020V110 CAVOK 21/16 Q1025 NOSIG
TAF LFTW 070800Z 0709/0809 VRB04KT CAVOK TEMPO 0801/0809 RA SCT014 OVC060

Avignon-Caumont Airport
LFMV 070930Z AUTO VRB02KT 9999 FEW062 20/15 Q1025 NOSIG
TAF LFMV 070800Z 0709/0809 VRB03KT CAVOK PROB30 TEMPO 0805/0809 RA FEW014 OVC080

Aérodrome de Carpentras
-
-

```

LFNG (Candillargues) and LFNH (Carpentras) are small aerodrome without weather station, in this case the messages are replaced with a simple dash (`-`).

### Options
- `--apiKey` (`-k`) mandatory, your CheckWX API key.
- `--icaoCodes` (`-c`) mandatory, a list of ICAO airport codes (4 letters) separated by a comma.

## License
[MIT](./LICENSE)
