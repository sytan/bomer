<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The three meta tag must be at the first beginning! -->
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Dashboard</title>

    <!-- Bootstrap core CSS -->
    <link href="../static/bootstrap-3.3.5/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="../static/css/dashboard.css" rel="stylesheet">
  </head>

  <body>

    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container-fluid">
        <!-- Brand and toggle get grouped for better mobile display -->
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar-collapse" aria-expanded="false">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>

          <a class="navbar-brand btn" role="button" data-toggle="collapse" href="#sidebar-collapse" aria-expanded="false" aria-controls="collapseExample">
            <span class="glyphicon glyphicon-th-list" aria-hidden="true"></span>
          </a>
        </div>

        <!-- Collect the nav links, forms, and other content for toggling -->
        <div class="collapse navbar-collapse" id="navbar-collapse">
          <ul class="nav navbar-nav navbar-right">
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Setting<span class="caret"></span></a>
              <ul class="dropdown-menu">
                <li><a href="#">{{.username}}</a></li>
                <li><a href="#">Log Out</a></li>
                <!-- <li role="separator" class="divider"></li> -->
              </ul>
            </li>
          </ul>
          <form class="navbar-form navbar-right" role="search">
            <div class="form-group">
              <input type="text" class="form-control" placeholder="Search">
            </div>
            <a class="btn my-btn" role="button" href="">
              <span class="glyphicon glyphicon-search my-navbar-icon"></span>
            </a>
          </form>
        </div><!-- /.navbar-collapse -->
      </div><!-- /.container-fluid -->
    </nav>

    <div class="collapse my-sidebar" id="sidebar-collapse">
      <div class="well">
        <ul class="nav nav-pills nav-stacked">
          <li role="presentation">
            <form enctype="multipart/form-data" method="post">
              <input type="file" name="uploadFile">
              <input type="submit">
            </form>
          </li>
          <li role="presentation"><a href="#">Profile</a></li>
          <li role="presentation"><a href="#">Messages</a></li>
        </ul>
      </div>
    </div>

    <!-- container start -->
    <div class="my-form-startRow">
      <div class="my-table">
        <span>After selecting your columns, click the 'Submit' button to upload your file.</span>
        <form action="analyse" method="post">
          <table class="table table-striped">
            <tr>
              <th><select class="form-control my-form-control-index" name="startIndex"> <!-- my-form-control-index to limit width</-->
                <option value="1">Select where to start</option>
                {{range $index,$_ := .rows}}
                <option value="{{$index}}">{{add $index 1}}</option>
                {{end}}
              </select></th>
              {{range .maxLengthCells}}
    					<th><select class="form-control" name="tableHeader">
    						<option value="Select A Column">Select A Column</option>
    						<option value="Manufacturer Part Number">Manufacturer Part Number</option>
    						<option value="Manufacturer Name">Manufacturer Name</option>
    						<option value="Digi-Key Part Number">Digi-Key Part Number</option>
    						<option value="Customer Reference">Customer Reference</option>
    						<option value="Quantity 1">Quantity 1</option>
    						<option value="Quantity 2">Quantity 2</option>
    						<option value="Quantity 3">Quantity 3</option>
    					</select></th>
              {{end}}
            </tr>

            {{range $index,$row := .rows}}
              <tr>
                <td>{{add $index 1}}</td> <!--call add to increase $index-->
                {{range $_,$cell := $row.Cells}}
                <td>{{$cell.String}}</td>
                {{end}}
              </tr>
            {{end}}
          </table>
          <input type="submit" class="btn btn-default" value="Submit">
        </form>
      </div>
    </div><!--container end-->

    <footer>
      <div class="author">
        © 2016 :
        <a href="{{.website}}">{{.company}}</a>
         - All Rights Reserved. Terms of use. Design by
        <a class="email" href="mailto:{{.email}}">{{.author}}</a>
      </div>
    </footer>
    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="../static/js/jquery-3.1.0.js"></script>
    <script src="../static/bootstrap-3.3.5/js/bootstrap.min.js"></script>
  </body>
</html>
