package gui

import (
	"fmt"

	"github.com/Pow-Duck/GetChiaKey/internal/chia"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type GUI struct {
	*vcl.TForm
	mainMenu *vcl.TMainMenu
}

var (
	mainForm *GUI
)

func InitGUI() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.Run()
}

func (g *GUI) OnFormCreate(sender vcl.IObject) {

	g.SetCaption("GetChiaKey ChiaKey获取工具 By: PowDuck.com 算力鸭 QQ交流群: 816250346")

	notice := vcl.NewLabel(g)
	notice.SetCaption("No network, no data security, code open source transparency https://github.com/Pow-Duck")
	notice.SetBounds(100, 50, 120, 32)
	notice.SetParent(g)
	notice2 := vcl.NewLabel(g)
	notice2.SetCaption("全程无连网保证你的数据安全， 代码开源透明 https://github.com/Pow-Duck")
	notice2.SetBounds(100, 70, 120, 32)
	notice2.SetParent(g)

	g.SetWidth(1000)
	g.SetHeight(600)

	g.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
		*CanClose = vcl.MessageDlg("Whether to quit 是否退出?", types.MtInformation, types.MbYes, types.MbNo) == types.MrYes

		if *CanClose {
			fmt.Println("OnCloseQuery")
		}
	})

	// 居中
	g.ScreenCenter()

	// bz
	fingerprint := vcl.NewLabel(g)
	fingerprint.SetCaption("指纹(Fingerprint): ")
	fingerprint.SetBounds(100, 150, 32, 32)
	fingerprint.SetParent(g)

	masterPublicKey := vcl.NewLabel(g)
	masterPublicKey.SetCaption("MasterPublicKey(MasterPublicKey): ")
	masterPublicKey.SetBounds(100, 200, 32, 32)
	masterPublicKey.SetParent(g)

	farmerPublicKey := vcl.NewLabel(g)
	farmerPublicKey.SetCaption("农民PublicKey(FarmerPublicKey): ")
	farmerPublicKey.SetBounds(100, 250, 32, 32)
	farmerPublicKey.SetParent(g)

	poolPublicKey := vcl.NewLabel(g)
	poolPublicKey.SetCaption("池PublicKey(PoolPublicKey): ")
	poolPublicKey.SetBounds(100, 300, 32, 32)
	poolPublicKey.SetParent(g)

	firstWalletAddress := vcl.NewLabel(g)
	firstWalletAddress.SetCaption("钱包地址(FirstWalletAddress): ")
	firstWalletAddress.SetBounds(100, 350, 32, 32)
	firstWalletAddress.SetParent(g)

	fingerprintInput := vcl.NewEdit(g)
	fingerprintInput.SetParent(g)
	fingerprintInput.SetLeft(380)
	fingerprintInput.SetTop(150)
	fingerprintInput.SetWidth(600)

	masterPublicKeyInput := vcl.NewEdit(g)
	masterPublicKeyInput.SetParent(g)
	masterPublicKeyInput.SetLeft(380)
	masterPublicKeyInput.SetTop(200)
	masterPublicKeyInput.SetWidth(600)

	farmerPublicKeyInput := vcl.NewEdit(g)
	farmerPublicKeyInput.SetParent(g)
	farmerPublicKeyInput.SetLeft(380)
	farmerPublicKeyInput.SetTop(250)
	farmerPublicKeyInput.SetWidth(600)

	poolPublicKeyInput := vcl.NewEdit(g)
	poolPublicKeyInput.SetParent(g)
	poolPublicKeyInput.SetLeft(380)
	poolPublicKeyInput.SetTop(300)
	poolPublicKeyInput.SetWidth(600)

	firstWalletAddressInput := vcl.NewEdit(g)
	firstWalletAddressInput.SetParent(g)
	firstWalletAddressInput.SetLeft(380)
	firstWalletAddressInput.SetTop(350)
	firstWalletAddressInput.SetWidth(600)

	btn := vcl.NewButton(g)
	btn.SetParent(g)
	btn.SetBounds(230, 450, 180, 30)
	btn.SetCaption("获取KEY(Get Key)")
	btn.SetOnClick(func(sender vcl.IObject) {
		key, err := chia.GetKey()
		if err != nil {
			vcl.ShowMessage("Get failed 获取失败 请联系 QQ交流群: 816250346")
			return
		}

		fingerprintInput.SetText(key.Fingerprint)
		masterPublicKeyInput.SetText(key.MasterPublicKey)
		farmerPublicKeyInput.SetText(key.FarmerPublicKey)
		poolPublicKeyInput.SetText(key.PoolPublicKey)
		firstWalletAddressInput.SetText(key.FirstWalletAddress)
	})
	// menu
	g.mainMenuInit()
}

func (g *GUI) mainMenuInit() {
	menu := vcl.NewMainMenu(g)

	aboutList := vcl.NewMenuItem(g)
	aboutList.SetCaption("About(&F)")
	menu.Items().Add(aboutList)

	aboutButton := vcl.NewMenuItem(g)
	aboutButton.SetCaption("About")
	aboutButton.SetOnClick(func(vcl.IObject) {
		vcl.ShowMessage("By: PowDuck.com 2021 QQ交流群: 816250346")
	})
	aboutList.Add(aboutButton)

	exitButton := vcl.NewMenuItem(g)
	exitButton.SetCaption("Exit(&E)")
	exitButton.SetShortCutFromString("Ctrl+Q")
	exitButton.SetOnClick(func(vcl.IObject) {
		mainForm.Close()
	})
	aboutList.Add(exitButton)

	l2 := vcl.NewLabel(g)
	l2.SetCaption(fmt.Sprintf("current version 当前版本: %s  QQ交流群: 816250346", "0.0.1"))
	l2.SetAlign(types.AlBottom)
	l2.SetParent(g)
}