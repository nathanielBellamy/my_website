<script lang="ts">
  import { FooterIcon } from "flowbite-svelte"
  import githubLogo from '../assets/github_logo.png'
  import linkedInLogo from '../assets/linked_in_logo.png'
  import gmailLogo from '../assets/gmail_logo.png'
</script>

<div class="h-full w-fit flex justify-between items-center gap-4">
  <FooterIcon href="https://github.com/nathanielBellamy" 
              class="text-gray-400 hover:text-gray-900 h-10 flex justify-around items-center">
    <img src={githubLogo}
         title="GitHub"
         class="invert"
         style:width="30px"
         style:height="30px"
         alt="GitHub"/>
  </FooterIcon>
  <FooterIcon href="https://www.linkedin.com/in/nateschieber/" 
              class="text-gray-400 hover:text-gray-900 h-10 flex justify-around items-center">
    <img src={linkedInLogo}
         title="LinkedIn"
         style:width="30px"
         style:height="30px"
         alt="LinkedIn"/>
  </FooterIcon>
  <FooterIcon href="mailto:nbschieber@gmail.com" 
              class="w-fit text-gray-400 hover:text-gray-900 h-10 flex justify-around items-center">
    <img src={gmailLogo}
         title="gmail"
         class="invert"
         style:width="30px"
         style:height="30px"
         alt="gmail"/>
  </FooterIcon>
</div>
