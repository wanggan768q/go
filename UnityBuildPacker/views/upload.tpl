<html>
    <head>
        <title>上传技能表</title>
    </head>
	
    <body>
        <form id="uploadForm" method="POST" enctype="multipart/form-data" action="/Upload/File">
            <p>上传技能表</p>
            <input type="FILE" accept=".xlsx" id="file" name="file" />
            <input type="SUBMIT" value="上传">

        </form>
        
        
        <script type="text/javascript">
            function Succeed() {
                alert("成功");
            }
            function Failure() {
                alert("失败");
            }
        </script>

        {{if .IsSucceed}}
        <script type="text/javascript">Succeed()</script>
        {{end}}
		{{if .IsFailure}}
        <script type="text/javascript">Failure()</script>
        {{end}}
    </body>
</html>