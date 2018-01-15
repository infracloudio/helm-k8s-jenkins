package main

const (
	html = `<HTML>
<HEAD>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/css/materialize.min.css">
	<TITLE>Go Lang Test App</TITLE>
	<STYLE TYPE="text/css">
	<!--
		@page { margin: 2cm }
		P { margin-bottom: 0.21cm }
		TD P { margin-bottom: 0cm }
	-->
	</STYLE>
</HEAD>
<BODY LANG="en-IN" DIR="LTR">
<div class="container">
<div class="row">
<div class="col s2">&nbsp;</div>
<div class="col s8">
<div class="card blue">
<div class="card-content white-text">
<div class="card-title">Pod that serviced this request</div>
</div>
<div class="card-content white">
<TABLE CLASS="bordered">
<TBODY>
	<TR>
		<TD>Pod ID</TD>
                <TD>{{.Pod}}</TD>
	</TR>
	<TR>
	        <TD>Version</TD>
                <TD>{{.Version}}</TD>
	</TR>
	<TR>
                <TD>Timestamp</TD>
                <TD>{{.TS}}</TD>
	</TR>
</TBODY>
</TABLE>
<P STYLE="margin-bottom: 0cm"><BR>
</P>
</BODY>
</HTML>`
)

