package engine

func (engine *Engine) NumTokenIndexAdded() uint64 {
	return engine.numTokenIndexAdded
}

func (engine *Engine) NumDocumentsIndexed() uint64 {
	return engine.numDocumentsIndexed
}
