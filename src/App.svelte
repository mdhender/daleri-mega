<script>
  let site = {
    title: 'Daleri Mega v1.2',
    slug: 'There is always room for more...',
    copyright: {year: 2019, author: 'Michael D Henderson'},
  };
  let mainMenu = [
    {id: 1, link: "#", label: "Main pages", class: "nav", children: [
      {id: 1.1, link: "#", label: "Blog"},
      {id: 1.2, link: "#", label: "Archives"},
      {id: 1.3, link: "#", label: "Categories"},
    ]},
    {id: 2, link: "#", label: "Blog topics", class: "nav", children: [
      {id: 2.1, link: "#", label: "Web design"},
      {id: 2.2, link: "#", label: "Accessibility"},
      {id: 2.3, link: "#", label: "CMS solutions"},
    ]},
    {id: 3, link: "#", label: "Extras", class: "nav", children: [
      {id: 3.1, link: "#", label: "Music archive"},
      {id: 3.2, link: "#", label: "Photo gallery"},
      {id: 3.3, link: "#", label: "Poems and lyrics"},
    ]},
    {id: 4, link: "#", label: "Community", class: "nav", children: [
      {id: 4.1, link: "#", label: "Guestbook"},
      {id: 4.2, link: "#", label: "Members"},
      {id: 4.3, link: "#", label: "Link collection"},
    ]},
    {id: 5, link: "#", label: "Introduction", class: "introduction", children: [
      {id: 5.1, label: "Space for a short presentation, a small banner or something else that could need some extra attention."},
    ]},
  ];
  let leftMenuLinks = [
    {id: 1, link: "#", label: "First page"},
    {id: 2, link: "#", label: "Second page"},
    {id: 3, link: "#", label: "Third page with subs", children: [
      {id: 3.1, link: "#", label: "First subpage"},
      {id: 3.2, link: "#", label: "Second subpage"},
    ]},
    {id: 4, link: "#", label: "Fourth page"},
  ];
  let rightMenuLinks = [
    {id: 1, link: "#", label: "Sixth page"},
    {id: 2, link: "#", label: "Seventh page"},
    {id: 3, link: "#", label: "Another page"},
    {id: 4, link: "#", label: "The last one"},
  ];
  let sampleLinks = [
    {id: 1, link: "#", label: "Sample link 1"},
    {id: 2, link: "#", label: "Sample link 2"},
    {id: 3, link: "#", label: "Sample link 3"},
    {id: 4, link: "#", label: "Sample link 4"},
  ];
</script>

<div id="top">
  <p id="skiplinks">Skip to: <a href="#content">content</a> | <a href="#sidebar">sidebar</a></p>
  <div id="sitetitle">
    <h1><a href="#">{site.title}</a></h1>
    <p>{site.slug}</p>
  </div>
  <hr class="clear" />
</div>

<div id="wrap">
  <div id="mainmenu">
    {#each mainMenu as menu}
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
    <h2><a href="#">Introducing: Daleri Mega</a></h2>
    <p class="introtext">Daleri Mega is a free, high-quality, standards-compliant XHTML/CSS template by Swedish template designer Andreas Viklund. It can be downloaded and used by anyone, and there are no limitations or obligations to worry about.</p>
    <p>Daleri Mega was built to offer lots of options for websites that will contain a lot of content. Through a set of various styles, you can arrange the layout in 2-5 columns. By default, the layout and all columns will follow the width of the browser window. The layout should work on any regular width (from 800px and up) and in all modern browsers including text-based browsers and handheld devices such as mobile phones. The main navigation is created with definition lists, and the menu can be expanded to contain more links if needed. You can also use the main menu as an overview or a basic sitemap, and expand the site with subpages using the sidebar buttons or regular link lists. On sites with fewer pages, the definition list menu can be removed completely, or replaced with an ad banner or a large header image. Feel free to experiment with the spaces to find your own preferred layout.</p>

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
    <hr class="clear" />
  </div>

  <div id="sidebar">
    <h2>Your own Mega...</h2>
    <p>Space for presentation, advertisements or extras - or maybe additional navigation links? You decide!</p>

    <div class="left">
      <h2>Left menu</h2>
      <ul class="sidemenu">
        {#each leftMenuLinks as link}
          <li>
            <a href="{link.link}">{link.label}</a>
            {#if link.children}
              <ul>
                {#each link.children as child}
                  <li><a href="{child.link}">{child.label}</a></li>
                {/each}
              </ul>
            {/if}
          </li>
        {/each}
      </ul>

      <h2>Notice</h2>
      <p>For more information on how to use this template to build your own website, and for tools and tutorials that can be useful, go to <a href="http://andreasviklund.com/">andreasviklund.com</a>.</p>
    </div>

    <div class="right">
      <h2>Right menu</h2>
      <ul class="sidemenu">
        {#each rightMenuLinks as { id, link, label }}
          <li><a href="{link}">{label}</a></li>
        {/each}
      </ul>

      <h2>Sample links</h2>
      <ul>
        {#each sampleLinks as { id, link, label }}
          <li><a href="{link}">{label}</a></li>
        {/each}
      </ul>
    </div>
    <hr class="clear" />
  </div>
  <hr class="clear" />
</div>

<div id="footer">
  <div class="left">
    <p>&copy; {site.copyright.year} <a href="#">{site.copyright.author}</a> | Template design by <a href="http://andreasviklund.com/">Andreas Viklund</a> | Best hosted at <a href="https://www.svenskadomaner.se/?ref=mall&amp;ling=en" title="Svenska Domäner AB">www.svenskadomaner.se</a></p>
  </div>
  <div class="right textright">
    <p>Sample footer menu: <a href="#">Link 1</a> | <a href="#">Link 2</a> | <a href="#">Link 3</a></p>
    <p class="hide"><a href="#top">Return to top</a></p>
  </div>
</div>
