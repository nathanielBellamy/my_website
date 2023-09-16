<script lang="ts">
  import { FooterIcon } from "flowbite-svelte"
  import Icon from "./Icon.svelte"
  import { Icons } from "./Icons"
</script>

<div class="h-full w-fit flex justify-between items-center gap-4">
  <FooterIcon href="https://github.com/nathanielBellamy" 
              class="text-cyan-700 hover:text-cyan-900 h-10 flex justify-around items-center">
    <Icon icon={Icons.GithubSolid}/>
  </FooterIcon>
  <FooterIcon href="https://www.linkedin.com/in/nateschieber/" 
              class="text-cyan-700 hover:text-cyan-900 h-10 flex justify-around items-center">
    <Icon icon={Icons.LinkedinSolid}/>
  </FooterIcon>
  <FooterIcon href="mailto:nbschieber@gmail.com" 
              class="text-cyan-700 hover:text-cyan-900 h-10 flex justify-around items-center">
    <Icon icon={Icons.MailBoxSolid}/>
  </FooterIcon>
</div>
