# Golang小工具
### Json解析工具（自用，实际上有个叫gjson的包已经很好用了）

```
		j, err := lahee.New(`{"ruleMap":{"short":{"rules":[{"key":"k1","range":{"[3-5]":3,"[6-9]":5,">=10":7},"decline":{"[604800-2592000)":0.8,"[2592000-5184000)":0.6,">=5184000":0.4},"weight":2},{"key":"k2","range":{"[3-5]":3,"[6-9]":5,">=10":7},"weight":2},{"key":"k3","range":{"[3-5]":3,"[6-9]":5,">=10":7},"weight":1.5},{"key":"k4","range":{"[3-5]":3,"[6-9]":5,">=10":7},"weight":1.5},{"key":"k5","range":{"[3-5]":3,"[6-9]":5,">=10":7},"decline":{"[604800-1209600)":0.8,"[1209600-1814400)":0.6,"[1814400-2419200)":0.4,">=2419200":0.2},"weight":1},{"key":"k6","range":{"[3-5]":3,"[6-9]":5,">=10":7},"weight":0.5}],"normalize":true}}}`)

	if err != nil {
		fmt.Println(err)
	} else {
		if shortList := j.PathKey("ruleMap").GetKey("long").GetKey("rules"); shortList.IsArrNode() && shortList.Size() != 0 {
			for i := 0; i < shortList.Size(); i++ {
				fmt.Println(shortList.GetIndex(i).GetKey("key").StringOrDefault(""))
			}
		}
	}

```
