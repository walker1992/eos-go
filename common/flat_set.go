package common

type Element interface {
	GetKey() uint64
}

type FlatSet struct {
	Data []Element
}

func (f *FlatSet) Len() int {
	return len(f.Data)
}

func (f *FlatSet) GetData(i int) Element {
	if len(f.Data)-1 >= i {
		return f.Data[i]
	}
	return nil
}

func (f *FlatSet) Clear() {
	if len(f.Data) > 0 {
		f.Data = nil
	}
}

func (f *FlatSet) Insert(element Element) (Element, bool) {
	var result Element
	length := f.Len()
	target := f.Data
	exist := false
	r := 0
	if length == 0 {
		f.Data = append(f.Data, element)
		result = f.Data[0]
	} else {
		i, j := 0, length-1
		for i < j {
			h := int(uint(i+j) >> 1)
			if target[h].GetKey() <= element.GetKey() {
				i = h + 1
			} else {
				j = h
			}
			r = h
		}
		if i <= j {
			if i == 0 || i == length-1 {
				//insert target before
				if element.GetKey() <= target[0].GetKey() {
					elemnts := []Element{}
					elemnts = append(elemnts, element)
					elemnts = append(elemnts, target...)
					f.Data = elemnts
					result = elemnts[0]
				} else if element.GetKey() >= target[length-1].GetKey() { //target append
					target = append(target, element)
					result = target[length]
					f.Data = target
				}
			} else {
				//Insert middle
				if element.GetKey() == target[r].GetKey() {
					element = target[r]
					result = target[r]
					exist = true
				} else {
					elemnts := []Element{}
					first := target[:r]
					second := target[r:length]
					elemnts = append(elemnts, first...)
					elemnts = append(elemnts, element)
					elemnts = append(elemnts, second...)
					f.Data = elemnts
					result = elemnts[r]
				}
			}
		}
	}
	return result, exist
}