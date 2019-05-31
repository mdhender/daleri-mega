import { writable, derived } from 'svelte/store';

const getInitialState = () => ({
    site: {
        title: 'Daleri Mega v1.2',
        slug: 'There is always room for more...',
        copyright: { year: 2019, author: 'Michael D Henderson' },
    },
    mainMenu: [
        {
            id: 1, link: "#", label: "Main pages", class: "nav", children: [
                { id: 1.1, link: "#", label: "Blog" },
                { id: 1.2, link: "#", label: "Archives" },
                { id: 1.3, link: "#", label: "Categories" },
            ]
        },
        {
            id: 2, link: "#", label: "Blog topics", class: "nav", children: [
                { id: 2.1, link: "#", label: "Web design" },
                { id: 2.2, link: "#", label: "Accessibility" },
                { id: 2.3, link: "#", label: "CMS solutions" },
            ]
        },
        {
            id: 3, link: "#", label: "Extras", class: "nav", children: [
                { id: 3.1, link: "#", label: "Music archive" },
                { id: 3.2, link: "#", label: "Photo gallery" },
                { id: 3.3, link: "#", label: "Poems and lyrics" },
            ]
        },
        {
            id: 4, link: "#", label: "Community", class: "nav", children: [
                { id: 4.1, link: "#", label: "Guestbook" },
                { id: 4.2, link: "#", label: "Members" },
                { id: 4.3, link: "#", label: "Link collection" },
            ]
        },
        {
            id: 5, link: "#", label: "Introduction", class: "introduction", children: [
                { id: 5.1, label: "Space for a short presentation, a small banner or something else that could need some extra attention." },
            ]
        },
    ],
    menu: {
        left: [
            { id: 1, link: "#", label: "First page" },
            { id: 2, link: "#", label: "Second page" },
            {
                id: 3, link: "#", label: "Third page with subs", children: [
                    { id: 3.1, link: "#", label: "First subpage" },
                    { id: 3.2, link: "#", label: "Second subpage" },
                ]
            },
            { id: 4, link: "#", label: "Fourth page" },
        ],
        right: [
            { id: 1, link: "#", label: "Sixth page" },
            { id: 2, link: "#", label: "Seventh page" },
            { id: 3, link: "#", label: "Another page" },
            { id: 4, link: "#", label: "The last one" },
        ],
        sample: [
            { id: 1, link: "#", label: "Sample link 1" },
            { id: 2, link: "#", label: "Sample link 2" },
            { id: 3, link: "#", label: "Sample link 3" },
            { id: 4, link: "#", label: "Sample link 4" },
        ],
    },
    notice: {
        title: 'Notice',
        text: 'For more information on how to use this template to build your own website, and for tools and tutorials that can be useful, go to <a href="http://andreasviklund.com/">andreasviklund.com</a>.',
    },
    content: {
        title: 'Introducing: Daleri Mega',
        paragraphs: [
            'Daleri Mega is a free, high-quality, standards-compliant XHTML/CSS template by Swedish template designer Andreas Viklund. It can be downloaded and used by anyone, and there are no limitations or obligations to worry about.',
            'Daleri Mega was built to offer lots of options for websites that will contain a lot of content. Through a set of various styles, you can arrange the layout in 2-5 columns. By default, the layout and all columns will follow the width of the browser window. The layout should work on any regular width (from 800px and up) and in all modern browsers including text-based browsers and handheld devices such as mobile phones. The main navigation is created with definition lists, and the menu can be expanded to contain more links if needed. You can also use the main menu as an overview or a basic sitemap, and expand the site with subpages using the sidebar buttons or regular link lists. On sites with fewer pages, the definition list menu can be removed completely, or replaced with an ad banner or a large header image. Feel free to experiment with the spaces to find your own preferred layout.',
        ],
    },
});

export const site = writable(getInitialState());