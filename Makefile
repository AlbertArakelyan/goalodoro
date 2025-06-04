BINARY_NAME_MACOS=Goalodoro.app

BINARY_NAME_WINDOWS=Goalodoro.exe
APP_ID_WINDOWS=com.goalodoro.aa

APP_NAME=Goalodoro
VERSION=0.1.1
BUILD_NO=1

build-macos:
	rm -rf ${BINARY_NAME_MACOS}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release -icon Icon.png

build-windows:
	rm ${BINARY_NAME_WINDOWS}
	fyne package -os windows -name ${BINARY_NAME_WINDOWS} -appID ${APP_ID_WINDOWS} -release