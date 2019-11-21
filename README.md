
MarkdownファイルをTeXファイルに変換し，テンプレートに挿入する．

# Usage

## テンプレートの作成
```
$mkdir template
$cat << EOF > template/template.tex
\documentclass[11pt]{jarticle}
\usepackage{ascmac}
\usepackage[dvipdfmx]{graphicx}
\usepackage[dvipdfmx]{hyperref}
\usepackage{url}
\usepackage{listings, jlisting}
\usepackage{here,txfonts,txfonts}
\usepackage{pxjahyper}
\usepackage{color}

\def\tightlist{\itemsep1pt\parskip0pt\parsep0pt}

\title{}
\author{}
\begin{document}
\maketitle

{{.TEXT}}

\end{document}
EOF
```

## templateファイルのコンパイル

```
$ make deps
$ statik -src=template
```

## 実行

### ローカルマシン

```
$ go build -o mt mt.go
$ mt markdown-file
```

### Docker

```
$ docker build -t xxx/mt .
$ docker run --rm -it -v $PWD:/work xxx/mt markdown-file
```
