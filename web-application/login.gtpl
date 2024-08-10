<html>
  <head>
    <title></title>
  </head>
  <body>
    <form action="/login" method="post">
      Username:<input type="text" name="username">
      <br>

      Password:<input type="password" name="password">
      <br>

      <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banana">banana</option>
      </select>
      <br>

      <input type="radio" name="gender" value="1">Male
      <input type="radio" name="gender" value="2">Female
      <br>

      <input type="checkbox" name="interest" value="football">Football
      <input type="checkbox" name="interest" value="basketball">Basketball
      <input type="checkbox" name="interest" value="tennis">Tennis
      <br>
      
      <input type="hidden" name="token" value="{{.}}">
      <input type="submit" value="Login">
    </form>
  </body>
</html>