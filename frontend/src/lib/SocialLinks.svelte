<script lang="ts">
  import { FooterIcon } from "flowbite-svelte"
  import Icon from "./Icon.svelte"
  import { Icons } from "./Icons"
</script>

<div class="h-full w-fit flex justify-between items-center gap-4"
     data-testid="social_links">
  <FooterIcon href="https://github.com/nathanielBellamy" 
              class="text-cyan-700 hover:text-cyan-900 h-10 flex justify-around items-center"
              data-testid="github_social_icon">
    <Icon icon={Icons.GithubSolid}/>
  </FooterIcon>
  <FooterIcon href="https://www.linkedin.com/in/nateschieber/" 
              class="text-cyan-700 hover:text-cyan-900 h-10 flex justify-around items-center"
              data-testid="linkedin_social_icon">
    <Icon icon={Icons.LinkedinSolid}/>
  </FooterIcon>
  <FooterIcon href="mailto:nbschieber@gmail.com" 
              class="text-cyan-700 hover:text-cyan-900 h-10 flex justify-around items-center"
              data-testid="mailto_social_icon">
    <Icon icon={Icons.MailBoxSolid}/>
  </FooterIcon>
</div>
