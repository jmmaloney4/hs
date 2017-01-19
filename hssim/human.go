package hssim

type HumanPlayer struct {
}

func (player HumanPlayer) InputType() InputType {
    return InputTypeCommandLine
}

func NewHumanPlayer() *HumanPlayer {
    rv := new(HumanPlayer)
    return rv
}
