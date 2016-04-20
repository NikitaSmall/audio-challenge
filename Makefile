BIN_DIR=bin
BYNARY=audio-challenge

BYNARY_OUTPUT= ${BIN_DIR}/${BYNARY}

all:
	go build -o ${BYNARY_OUTPUT}
