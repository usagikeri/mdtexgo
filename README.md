
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
$statik -src=template
$go build mt.go
```

## goファイルのbuild

```
$go build mt.go
```
