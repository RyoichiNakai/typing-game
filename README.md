# Typing Game
## 実装仕様
- 標準出力に英単語を出す（出すものは自由）
- 標準入力から1行受け取る 
- 制限時間内に何問解けたか表示する

>@tenntennさんが作成したGopher道場の課題です。 Gopher道場については以下のURLから参照してください。
> - https://gopherdojo.org/

## main.goの構成
- for-selectパターンを使用
    - mainとgoroutineをそれぞれループさせる
    - チャネルが渡されたときにそのチャネルと条件が合うものを実行
- 指定された秒数の計測と文字列を読み込む２種類のgoroutineを作成
    - context.WithTimeout
        - 指定された時間を計測し、チャネルとキャンセル処理を返す
    - input
        - 読み込んだ文字列をチャネルとして返す

## 参考URL
>https://kenzo0107.github.io/2020/04/29/2020-04-30-typing-game-go/
