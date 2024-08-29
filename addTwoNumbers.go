package main

import "fmt"

// 定義鏈表節點
type ListNode struct {
    Val  int
    Next *ListNode
}

// addTwoNumbers 函數實現兩個鏈表的數字相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}  // 虛擬頭節點
    current := dummy      // 當前節點指針
    carry := 0            // 進位變量

    // 遍歷兩個鏈表，直到兩個鏈表都為空
    for l1 != nil || l2 != nil {
        x := 0
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }

        y := 0
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }

        sum := carry + x + y   // 計算當前位的總和
        carry = sum / 10       // 更新進位，取整
        current.Next = &ListNode{Val: sum % 10} // 創建新節點保存當前位結果
        current = current.Next // 移動指針到下一個節點
    }

    // 處理最後的進位
    if carry > 0 {
        current.Next = &ListNode{Val: carry}
    }

    return dummy.Next // 返回結果鏈表的頭節點
}

// 輔助函數：將數組轉換為鏈表
func arrayToList(arr []int) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for _, v := range arr {
        current.Next = &ListNode{Val: v}
        current = current.Next
    }
    return dummy.Next
}

// 輔助函數：將鏈表轉換為數組，便於打印結果
func listToArray(list *ListNode) []int {
    var result []int
    for list != nil {
        result = append(result, list.Val)
        list = list.Next
    }
    return result
}

func main() {
    l1 := arrayToList([]int{2, 4, 3})
    l2 := arrayToList([]int{5, 6, 4})
    result := addTwoNumbers(l1, l2)
    fmt.Println(listToArray(result)) // 輸出 [7 0 8]

    l1 = arrayToList([]int{9, 9, 9, 9, 9, 9, 9})
    l2 = arrayToList([]int{9, 9, 9, 9})
    result = addTwoNumbers(l1, l2)
    fmt.Println(listToArray(result)) // 輸出 [8 9 9 9 0 0 0 1]
}
