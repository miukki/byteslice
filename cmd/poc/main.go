package main

import (
	"fmt"
	"os"
	"test-zkp/cmd/poc/byteslice"
	"test-zkp/cmd/poc/server"
	"test-zkp/cmd/poc/sha3"
	"test-zkp/cmd/poc/signing"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	GIT_COMMIT string
)

var (
	deployFlag = cli.BoolFlag{
		Name: "deploy",
	}
)

var (
	logFlag = cli.BoolFlag{
		Name: "loghash",
	}
)

var (
	gethashFlag = cli.BoolFlag{
		Name: "gethash",
	}
)

var (
	getlogFlag = cli.BoolFlag{
		Name: "getlog",
	}
)

var (
	gethistoryFlag = cli.BoolFlag{
		Name: "gethistory",
	}
)

var (
	verifylatestFlag = cli.BoolFlag{
		Name: "verifylatest",
	}
)

var (
	smiloFlag = cli.BoolFlag{
		Name: "smilo",
	}
)

var (
	verifysigFlag = cli.BoolFlag{
		Name: "verifysig",
	}
)

var (
	verifyfilesigFlag = cli.BoolFlag{
		Name: "verifyfilesig",
	}
)

var (
	texthashFlag = cli.BoolFlag{
		Name: "texthash",
	}
)

var (
	filehashFlag = cli.BoolFlag{
		Name: "filehash",
	}
)

var (
	filesignFlag = cli.BoolFlag{
		Name: "filesign",
	}
)

var (
	apikeyFlag = cli.BoolFlag{
		Name: "apikey",
	}
)

var (
	addressFlag = cli.StringFlag{
		Name:  "address",
		Usage: "--address=0x5dcCd00a71Ee168254e3c54E7d7bC43ebbba3E62",
	}
)

var (
	fileFlag = cli.StringFlag{
		Name:  "file",
		Usage: "--file=PoseidonPPT_Khovratovich.pdf",
	}
)

var (
	pubkeyFlag = cli.StringFlag{
		Name:  "pubkey",
		Usage: "--pubkey=0x222de25f8912859d498e861dC6241CA1c9C086E9",
	}
)

var (
	signatureFlag = cli.StringFlag{
		Name:  "signature",
		Usage: "--signature=0x9db472d772d1e213ef069077aa6ca3100f12a0a5255dbda6203307cdf782354b76ea9c8cf06583b1e39d439fd0b712b7a0862c1f9293cc02b5591e4e2f4b9ae300",
	}
)

var (
	txnFlag = cli.StringFlag{
		Name:  "txn",
		Usage: "--txn='{\"id\": \"1234567\",\"recipients\": [\"1000000\",\"111111\",\"999999\"]}'",
	}
)

var (
	msgFlag = cli.StringFlag{
		Name:  "msg",
		Usage: "--msg='Some message that we would like to sign'",
	}
)

var (
	txtFlag = cli.StringFlag{
		Name:  "txt",
		Usage: "--txt='Some text that we want to calculate a hash for'",
	}
)

var (
	zkpkeyFlag = cli.StringFlag{
		Name:  "zkpkey",
		Usage: "--zkpkey=Dvnq4vle3pxlnkERNXQyMh50QgBWKAHU7xel4NMsFYo=",
	}
)

var (
	typeFlag = cli.StringFlag{
		Name:  "type",
		Usage: "--type=OTCORDER",
	}
)

var (
	hashtypeFlag = cli.StringFlag{
		Name:  "hashtype",
		Usage: "--hashtype=NEWLEGACYKECCAK256",
	}
)

var (
	stateFlag = cli.StringFlag{
		Name:  "state",
		Usage: "--state=CREATED",
	}
)

var (
	idFlag = cli.StringFlag{
		Name:  "id",
		Usage: "--id=111111",
	}
)

var (
	debugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "debug",
	}
)

var (
	initFlag = cli.BoolFlag{
		Name:  "init",
		Usage: "init",
	}
)

var (
	chainFlag = cli.BoolFlag{
		Name:  "chain",
		Usage: "chain",
	}
)

func prepare() *cli.App {
	var app = cli.NewApp()

	app.Name = fmt.Sprintf("Test ZKP POC CMD - GIT_COMMIT=%s", GIT_COMMIT)
	app.Usage = "Test ZKP POC command line interface"

	//nolint
	app.Commands = []cli.Command{
		portsCMD,
		bytesliceCMD,
		serverCMD,
		signingCMD,
	}
	return app
}

func main() {
	app := prepare()
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	signingCMD = cli.Command{
		Name:      "signing",
		Usage:     "run the signing poc",
		Action:    signerAction,
		ArgsUsage: "",
		Flags: []cli.Flag{
			smiloFlag,
			zkpkeyFlag,
			pubkeyFlag,
			fileFlag,
			filesignFlag,
			signatureFlag,
			verifysigFlag,
			verifyfilesigFlag,
			msgFlag,
		},
		Description: "",
	}

	serverCMD = cli.Command{
		Name:        "server",
		Usage:       "run the backend server for this poc",
		Action:      serverAction,
		ArgsUsage:   "",
		Flags:       []cli.Flag{},
		Description: "",
	}

	bytesliceCMD = cli.Command{
		Name:      "byteslice",
		Usage:     "run the byteslice poc",
		Action:    bytesliceAction,
		ArgsUsage: "",
		Flags: []cli.Flag{
			initFlag,
			chainFlag,
			debugFlag,
		},
		Description: "",
	}

	portsCMD = cli.Command{
		Name:      "sha3",
		Usage:     "run the sha3 poc",
		Action:    sha3Action,
		ArgsUsage: "",
		Flags: []cli.Flag{
			deployFlag,
			logFlag,
			txnFlag,
			typeFlag,
			smiloFlag,
			apikeyFlag,
			zkpkeyFlag,
			stateFlag,
			addressFlag,
			texthashFlag,
			filehashFlag,
			fileFlag,
			hashtypeFlag,
			txtFlag,
			getlogFlag,
			gethashFlag,
			gethistoryFlag,
			verifylatestFlag,
			idFlag,
		},
		Description: "",
	}
)

func bytesliceAction(ctx *cli.Context) error {
	log.WithField("GIT_COMMIT", GIT_COMMIT).Info("bytesliceAction starting")

	initFlg := ctx.Bool(initFlag.Name)
	debugFlg := ctx.Bool(debugFlag.Name)
	chainFlg := ctx.Bool(chainFlag.Name)

	if initFlg {
		err := byteslice.Init()
		if err != nil {
			return err
		}
	} else if debugFlg {
		err := byteslice.Debug()
		if err != nil {
			return err
		}
	} else if chainFlg {
		err := byteslice.Chainverify()
		if err != nil {
			return err
		}
	} else {
		err := byteslice.Run()
		if err != nil {
			return err
		}
	}

	log.Info("bytesliceAction completed")
	return nil
}

func serverAction(ctx *cli.Context) error {
	log.WithField("GIT_COMMIT", GIT_COMMIT).Info("serverAction starting")
	router, err := server.Server("")
	if err != nil {
		return err
	}
	err = server.Start(router)
	if err != nil {
		return err
	}

	log.Info("serverAction completed")
	return nil
}

func signerAction(ctx *cli.Context) error {

	smiloFlg := ctx.Bool(smiloFlag.Name)
	verifysigFlg := ctx.Bool(verifysigFlag.Name)
	verifyfilesigFlg := ctx.Bool(verifyfilesigFlag.Name)
	filesignFlg := ctx.Bool(filesignFlag.Name)
	fileFlg := ctx.String(fileFlag.Name)
	signatureFlg := ctx.String(signatureFlag.Name)
	pubkeyFlg := ctx.String(pubkeyFlag.Name)
	zkpkeyFlg := ctx.String(zkpkeyFlag.Name)
	msgFlg := ctx.String(msgFlag.Name)

	if verifysigFlg {
		_, err := signing.VerifySignature(msgFlg, zkpkeyFlg, pubkeyFlg, signatureFlg)
		if err != nil {
			return err
		}
	} else if verifyfilesigFlg {
		_, err := signing.VerifyFileSignature(fileFlg, zkpkeyFlg, pubkeyFlg, signatureFlg)
		if err != nil {
			return err
		}
	} else {
		// First handle all smilo request, if the smilo flag was set, else assume working against a local Ganache
		if smiloFlg {
			if filesignFlg {
				_, err := signing.FileSigner(true, fileFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else {
				_, err := signing.Signer(true, msgFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			}
		} else {
			if filesignFlg {
				_, err := signing.FileSigner(false, fileFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else {
				_, err := signing.Signer(false, msgFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			}
		}
	}

	log.Info("signerAction completed")
	return nil
}

func sha3Action(ctx *cli.Context) error {
	log.WithField("GIT_COMMIT", GIT_COMMIT).Info("sha3Action starting")

	deployFlg := ctx.Bool(deployFlag.Name)
	logFlg := ctx.Bool(logFlag.Name)
	txnFlg := ctx.String(txnFlag.Name)
	typeFlg := ctx.String(typeFlag.Name)
	stateFlg := ctx.String(stateFlag.Name)
	addressFlg := ctx.String(addressFlag.Name)
	zkpkeyFlg := ctx.String(zkpkeyFlag.Name)
	getlogFlg := ctx.Bool(getlogFlag.Name)
	gethashFlg := ctx.Bool(gethashFlag.Name)
	gethistoryFlg := ctx.Bool(gethistoryFlag.Name)
	verifylatestFlg := ctx.Bool(verifylatestFlag.Name)
	smiloFlg := ctx.Bool(smiloFlag.Name)
	apikeyFlg := ctx.Bool(apikeyFlag.Name)
	idFlg := ctx.String(idFlag.Name)
	texthashFlg := ctx.Bool(texthashFlag.Name)
	txtFlg := ctx.String(txtFlag.Name)
	hashtypeFlg := ctx.String(hashtypeFlag.Name)
	filehashFlg := ctx.Bool(filehashFlag.Name)
	fileFlg := ctx.String(fileFlag.Name)

	// If a new API Key is requested, generate one and expect that to be used for every request
	if apikeyFlg {
		result, err := sha3.CreateApiKey()
		if err != nil {
			return err
		}
		log.WithField("Result", result).Info("This is the new API key, please update the environment")
	} else {

		// First handle all smilo request, if the smilo flag was set, else assume working against a local Ganache
		if smiloFlg {
			if deployFlg {
				_, err := sha3.Deploy(true, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if logFlg {
				_, _, err := sha3.Loghash(true, txnFlg, typeFlg, stateFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if getlogFlg {
				_, _, _, err := sha3.Getlog(true, idFlg, typeFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if gethashFlg {
				_, err := sha3.Gethash(true, idFlg, typeFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if gethistoryFlg {
				_, _, err := sha3.Gethistory(true, idFlg, typeFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if verifylatestFlg {
				_, err := sha3.VerifyLatest(true, idFlg, typeFlg, txnFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else {
				err := sha3.Run()
				if err != nil {
					return err
				}
			}
		} else {
			if deployFlg {
				_, err := sha3.Deploy(false, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if logFlg {
				_, _, err := sha3.Loghash(false, txnFlg, typeFlg, stateFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if getlogFlg {
				_, _, _, err := sha3.Getlog(false, idFlg, typeFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if gethashFlg {
				_, err := sha3.Gethash(false, idFlg, typeFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if texthashFlg {
				_, err := sha3.CalcTexthash(txtFlg, hashtypeFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if filehashFlg {
				_, err := sha3.CalcFilehash(fileFlg, hashtypeFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if gethistoryFlg {
				_, _, err := sha3.Gethistory(false, idFlg, typeFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
			} else if verifylatestFlg {
				result, err := sha3.VerifyLatest(false, idFlg, typeFlg, txnFlg, addressFlg, zkpkeyFlg)
				if err != nil {
					return err
				}
				log.WithField("Result", result).Info("Testing hash of json with latest state")
			} else {
				err := sha3.Run()
				if err != nil {
					return err
				}
			}
		}
	}

	log.Info("sha3Action completed")
	return nil
}
