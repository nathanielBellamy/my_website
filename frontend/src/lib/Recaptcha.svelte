<script lang="ts">
  export let hasPassed: boolean = false
  var onloadCallback = function() {
    grecaptcha.render('html_recaptcha_element', {
      'sitekey' : import.meta.env.VITE_GRECAPTCHA_KEY
    });
  };
</script>

<body>
  <!-- TODO -->
  <!--   - create Go endpoint and have this hit it -->
  <form action="?" method="POST">
    <div id="html_recaptcha_element"></div>
    <br>
    <input type="submit" value="Submit">
  </form>
  <script src="https://www.google.com/recaptcha/api.js?onload=onloadCallback&render=explicit"
      async defer>
  </script>
</body>

