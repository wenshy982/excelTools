# 按指定列拆excel使用说明：


打开excel表，

ALT+Fn+F11-->打开VBA编辑器-->插入-->模块

将下列代码拷贝至弹出的窗口

点击运行-->弹出窗口

选择需要拆分的列-->完成后显示done

---

新版代码  多了容易断
```VBA
Sub 拆分为多个excel文件()
 Dim arr, d As Object, k, t, i&, lc%, rng As Range, c%
 c = Application.InputBox("提示：请输入要拆分列号", , 3, , , , , 1)
tRow = Application.InputBox("提示：请您输入表的标题总行数？")
If tRow = 0 Then MsgBox "输入错误，程序将退出！": Exit Sub
Application.ScreenUpdating = False
Application.DisplayAlerts = False
arr = [a1].CurrentRegion
lc = UBound(arr, 2)
Set rng = [a1].Resize(tRow + 1, lc)
 Set d = CreateObject("scripting.dictionary")
For i = tRow + 1 To UBound(arr)
If Not d.Exists(arr(i, c)) Then
Set d(arr(i, c)) = Cells(i, 1).Resize(1, lc)
 Else
 Set d(arr(i, c)) = Union(d(arr(i, c)), Cells(i, 1).Resize(1, lc))
End If
 Next
k = d.Keys
 t = d.Items
tt = tRow + 1
For i = 0 To d.Count - 1
With Workbooks.Add(xlWBATWorksheet)
rng.Copy .Sheets(1).[a1]
 t(i).Copy .Sheets(1).[a4]
 .SaveAs Filename:=ThisWorkbook.Path & "\" & k(i) & ".xlsx"
 .Close
 End With
 Next
 Application.DisplayAlerts = True
 Application.ScreenUpdating = True
 MsgBox "拆分成功，欧耶"
End Sub
```
---

旧版代码

```VBA
Sub 保留表头拆分数据为若干新工作簿()
    Dim arr, d As Object, k, t, i&, lc%, rng As Range, c%
    c = Application.InputBox("请输入拆分列号", , 2, , , , , 1)
    If c = 0 Then Exit Sub
    Application.ScreenUpdating = False
    Application.DisplayAlerts = False
    arr = [a1].CurrentRegion
    lc = UBound(arr, 2)
    Set rng = [a1].Resize(, lc)
    Set d = CreateObject("scripting.dictionary")
    For i = 2 To UBound(arr)
        If Not d.Exists(arr(i, c)) Then
            Set d(arr(i, c)) = Cells(i, 1).Resize(1, lc)
        Else
            Set d(arr(i, c)) = Union(d(arr(i, c)), Cells(i, 1).Resize(1, lc))
        End If
    Next
    k = d.Keys
    t = d.Items
    For i = 0 To d.Count - 1
        With Workbooks.Add(xlWBATWorksheet)
            rng.Copy .Sheets(1).[a1]
            t(i).Copy .Sheets(1).[a2]
            .SaveAs Filename:=ThisWorkbook.Path & "\" & k(i) & ".xlsx"
            .Close
        End With
    Next
    Application.DisplayAlerts = True
    Application.ScreenUpdating = True
    MsgBox "完毕"
End Sub
```

