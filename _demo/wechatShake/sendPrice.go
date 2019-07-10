package main

func sendCoin(d *gift) (bool, string) {
	if d.total == 0 {
		return true, d.data
	} else if d.left > 0 {
		d.left--
		return true, d.data
	}
	return false, "No more"
}

func sendCoupon(d *gift) (bool, string) {
	if d.left > 0 {
		left := d.left - 1
		d.left = left

		if left >= len(d.dataList) {
			return false, "No more"
		}
		return true, d.dataList[left]
	}

	return false, "No more"
}

func sendCouponFix(d *gift) (bool, string) {
	if d.total == 0 {
		return true, d.data
	} else if d.left > 0 {
		d.left--
		return true, d.data
	}
	return false, "No more"
}

func sendRealLarge(d *gift) (bool, string) {
	if d.total == 0 {
		return true, d.data
	} else if d.left > 0 {
		d.left--
		return true, d.data
	}
	return false, "No more"
}

func sendRealSmall(d *gift) (bool, string) {
	if d.total == 0 {
		return true, d.data
	} else if d.left > 0 {
		d.left--
		return true, d.data
	}
	return false, "No more"
}

func saveLuckyData(code int32, id int, name string, link string, sendData string, left int) {
	logger.Printf("lucky, code: %d, gift: %d, name: %s, link: %s, data: %s, left: %d",
		code, id, name, link, sendData, left)
}
