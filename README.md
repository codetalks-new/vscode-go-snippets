# Go Snippets 
![Logo](https://github.com/banxi1988/vscode-go-snippets/blob/master/images/logo.png?raw=true)

Currently we only offer some snippets for leetcode . below is the snippets table:

| prefix | description |
| :----- | :------ |
| absi | abs for int |
| maxi | max for int |
| mini | min for int |
| maxints | max for ints |
| joinints | join ints to string |
| limits | MaxInt and MinInt |
| treenode | type TreeNode struct |
| dumptree | Dump Tree using level order |
| listnode | type ListNode struct |
| dumplistnode | Dump ListNode using level order |
| reversearr | Reverse Ints Array |
| reversestr | Reverse String |
| pbasic | Basic file for problem |
| pbinarytree | Binary Tree solution file stub |
| p1darr | One Dimen Array Basic file stub |
| p2darr | One Dimen Array Basic file stub |


And I have also add some slice tricks as snippet below:
taken from [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)

| prefix | description |
| :----- | :------ |
| range.map | range map |
| range.arr | range array, slices |
| range.str | range interate over unicode points,not bytes |
| slice.av | Slice Append Vector |
| slice.copy | Slice Copy |
| slice.cut | a = append(a[:i], a[j:]...) |
| slice.delete | a = append(a[:i], a[i+1:]...) |
| slice.extend | a = append(a, make([]T, j)...) |
| slice.insert | a = append(a[:i], append([]T{x}, a[i:]...)...) |
| slice.unshift | a = append([]T{x}, a...) |
| slice.push | a = append(a, x) |
| slice.pop | x, a = a[0], a[1:] |
| slice.popback | x, a = a[len(a) -1], a[:len(a) -1] |
