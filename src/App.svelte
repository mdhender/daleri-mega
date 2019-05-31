<script>
  import { site } from './stores/site';

  console.log(site);
</script>

<div id="top">
  <p id="skiplinks">Skip to: <a href="#content">content</a> | <a href="#sidebar">sidebar</a></p>
  <div id="sitetitle">
    <h1><a href="/">{$site.site.title}</a></h1>
    <p>{$site.site.slug}</p>
  </div>
  <hr class="clear" />
</div>

<div id="wrap">
  <div id="mainmenu">
    {#each $site.mainMenu as menu}
      <dl class="{menu.class}">
        <dt><a href="{menu.link}">{menu.label}</a></dt>
          {#each menu.children as link}
            {#if link.link}
              <dd><a href="{link.link}">{link.label}</a></dd>
            {:else}
              <dd>{link.label}</dd>
            {/if}
          {/each}
      </dl>
    {/each}
    <hr class="clear" />
  </div>

  <div id="content">
    {#if $site.content}
      <h2><a href="#">{$site.content.title}</a></h2>
      {#each $site.content.paragraphs as p, i}
        {#if i === 0}
          <p class="introtext">{p}</p>
        {:else}
          <p>{p}</p>
        {/if}
      {/each}

      <div class="trio1">
        <h3><a href="#">Support the development</a></h3>
        <p>If you want to support the template designer and encourage the development of more free templates, you should leave the footer credit link that says "Template design by Andreas Viklund". But if you need to remove the credits, you are free to do so. There are plenty of <a href="http://andreasviklund.com/support/">other ways</a> to support the designers work if you want to.</p>
      </div>

      <div class="trio2">
        <h3><a href="#">Sample text links</a></h3>
        <p>If you find this template useful, I recommend that you visit these sites for more useful material and resources:</p>
        <ul>
          <li><a href="http://andreasviklund.com/templates/">Free website templates</a></li>
          <li><a href="http://andreasviklund.com/wordpress-themes/">WordPress themes</a></li>
        </ul>
      </div>

      <div class="trio3">
        <h3><a href="#">Customization</a></h3>
        <p>With all the options, it can often become too much of the good stuff. This template allows you to remove parts that you don't need. For example, if you don't want the massive menu, you can remove it and just use link buttons in the sidebar. If you don't want all five columns, then just use two or three. Need professional help with the customization? Then <a href="http://andreasviklund.com/contact/">contact</a> the template designer for information about rates and availability!</p>
      </div>
    {/if}
    <hr class="clear" />
  </div>

  <div id="sidebar">
    <h2>Your own Mega...</h2>
    <p>Space for presentation, advertisements or extras - or maybe additional navigation links? You decide!</p>

    <div class="left">
      <h2>Left menu</h2>
      <ul class="sidemenu">
        {#each $site.menu.left as menu}
          <li>
            <a href="{menu.link}">{menu.label}</a>
            {#if menu.children}
              <ul>
                {#each menu.children as child}
                  <li><a href="{child.link}">{child.label}</a></li>
                {/each}
              </ul>
            {/if}
          </li>
        {/each}
      </ul>

      {#if $site.notice}
        <h2>{$site.notice.title}</h2>
        <p>{@html $site.notice.text}</p>
      {/if}
    </div>

    <div class="right">
      <h2>Right menu</h2>
      <ul class="sidemenu">
        {#each $site.menu.right as menu}
          <li><a href="{menu.link}">{menu.label}</a></li>
        {/each}
      </ul>

      <h2>Sample links</h2>
      <ul>
        {#each $site.menu.sample as menu}
          <li><a href="{menu.link}">{menu.label}</a></li>
        {/each}
      </ul>
    </div>
    <hr class="clear" />
  </div>
  <hr class="clear" />
</div>

<div id="footer">
  <div class="left">
    <p>&copy; {$site.site.copyright.year} <a href="#">{$site.site.copyright.author}</a> | Template design by <a href="http://andreasviklund.com/">Andreas Viklund</a> | Best hosted at <a href="https://www.svenskadomaner.se/?ref=mall&amp;ling=en" title="Svenska Domäner AB">www.svenskadomaner.se</a></p>
  </div>
  <div class="right textright">
    <p>Sample footer menu: <a href="#">Link 1</a> | <a href="#">Link 2</a> | <a href="#">Link 3</a></p>
    <p class="hide"><a href="#top">Return to top</a></p>
  </div>
</div>
