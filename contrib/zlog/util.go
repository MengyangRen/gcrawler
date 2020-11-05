package zlog

import "context"

const ctxDlogOrderedMapKey = "ctxElogOrderedMapKey"

func SetTraceInfo(ctx context.Context, traceID, pSpanID, spanID string) context.Context {
	om := NewOrderMap()
	om.Set("trace_id", traceID)
	om.Set("pspan_id", pSpanID)
	om.Set("span_id", spanID)
	src := FromContext(ctx)
	if src == nil {
		src = NewOrderMap()
	}
	src.AddVals(om)
	return setContext(ctx, src)
}
func GetTraceInfo(ctx context.Context) (traceID, pSpanID, spanID string) {
	om := FromContext(ctx)
	if tmp, ok := om.Get("trace_id"); ok {
		traceID = tmp.(string)
	}
	if tmp, ok := om.Get("pspan_id"); ok {
		traceID = tmp.(string)
	}
	if tmp, ok := om.Get("span_id"); ok {
		traceID = tmp.(string)
	}
	return
}

func FromContext(ctx context.Context) *OrderedMap {
	ret := ctx.Value(ctxDlogOrderedMapKey)
	if ret == nil {
		return nil
	}
	return ret.(*OrderedMap)
}

//别人不需要用
func setContext(ctx context.Context, dt *OrderedMap) context.Context {
	return context.WithValue(ctx, ctxDlogOrderedMapKey, dt)
}
