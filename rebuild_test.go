package etree

import (
	"fmt"
	//	"os"
	"testing"
	//"github.com/beevik/etree"
)

func TestCopy(t *testing.T) {
	s := `
  <head>
            <meta charset="utf-8"></meta>
            <t t-if="main_object and &apos;website_meta_title&apos; in main_object and not title">
                <t t-set="title" t-value="main_object.website_meta_title"></t>
            </t>
            <t t-if="main_object and &apos;name&apos; in main_object and not title and not additional_title">
                <t t-set="additional_title" t-value="main_object.name"></t>
            </t>
            <t t-if="not title">
                <t t-set="title"><t t-if="additional_title"><t t-raw="additional_title"></t> | </t><t t-esc="(website or res_company).name"></t></t>
            </t>

            <meta name="viewport" content="initial-scale=1"></meta>
            <meta name="description" t-att-content="main_object and &apos;website_meta_description&apos; in main_object
                and main_object.website_meta_description or website_meta_description"></meta>
            <meta name="keywords" t-att-content="main_object and &apos;website_meta_keywords&apos; in main_object
                and main_object.website_meta_keywords or website_meta_keywords"></meta>
            <meta name="generator" content="Odoo"></meta>

            <!-- OpenGraph tags for Facebook sharing -->
            <meta property="og:title" t-att-content="additional_title"></meta>
            <meta property="og:site_name" t-att-content="res_company.name"></meta>
            <t t-if="main_object and &apos;plain_content&apos; in main_object and main_object.plain_content">
                <t t-set="og_description" t-value="main_object.plain_content[0:500]"></t>
                <meta property="og:description" t-att-content="og_description"></meta>
                <meta property="og:image" t-att-content="request.httprequest.url_root+&apos;logo.png&apos;"></meta>
                <meta property="og:url" t-att-content="request.httprequest.url_root+request.httprequest.path[1:end]"></meta>
            </t>

            <title><t t-esc="title"></t></title>

            <t t-set="languages" t-value="website.get_languages() if website else None"></t>
            <t t-if="request and request.website_multilang and website">
                <t t-foreach="website.get_alternate_languages(request.httprequest)" t-as="lg">
                    <link rel="alternate" t-att-hreflang="lg[&apos;hreflang&apos;]" t-att-href="lg[&apos;href&apos;]"></link>
                </t>
            </t>

            <t t-call-assets="web.assets_common" t-js="false"></t>
            <t t-call-assets="website.assets_frontend" t-js="false"></t>
            <t t-call-assets="web.assets_common" t-css="false"></t>
            <t t-call-assets="website.assets_frontend" t-css="false"></t>
            <script type="text/javascript">
                odoo.define(&apos;web.csrf&apos;, function (require) {
                    var token = &quot;<t t-esc="request.csrf_token(None)"></t>&quot;;
                    require(&apos;web.core&apos;).csrf_token = token;
                    require(&apos;qweb&apos;).default_dict.csrf_token = token;
                });
            </script>

            <t t-raw="head or &apos;&apos;" name="layout_head"></t>
        </head>
        <body>
            <div id="wrapwrap">
                <header>
                    <div class="navbar navbar-default navbar-static-top">
                        <div class="container">
                            <div class="navbar-header">
                                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-top-collapse">
                                    <span class="sr-only">Toggle navigation</span>
                                    <span class="icon-bar"></span>
                                    <span class="icon-bar"></span>
                                    <span class="icon-bar"></span>
                                </button>
                                <a class="navbar-brand" href="/" t-if="website" t-field="website.name">My Website</a>
                            </div>
                            <div class="collapse navbar-collapse navbar-top-collapse">
                                <ul class="nav navbar-nav navbar-right" id="top_menu">
                                    <t t-foreach="website.menu_id.child_id" t-as="submenu">
                                        <t t-call="website.submenu"></t>
                                    </t>
                                    <li class="divider" t-ignore="true" t-if="website.user_id != user_id"></li>
                                    <li class="dropdown" t-ignore="true" t-if="website.user_id != user_id">
                                        <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                            <b>
                                                <span t-esc="(len(user_id.name)&gt;25) and (user_id.name[:23]+&apos;...&apos;) or user_id.name"></span>
                                                <span class="caret"></span>
                                            </b>
                                        </a>
                                        <ul class="dropdown-menu js_usermenu" role="menu">
                                            <li id="o_logout"><a t-attf-href="/web/session/logout?redirect=/" role="menuitem">Logout</a></li>
                                        </ul>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </header>
                <main>
                    <t t-raw="0"></t>
                </main>
                <footer>
                    <div id="footer">
                    </div>
                </footer>
            </div>
            <script id="tracking_code" t-if="website and website.google_analytics_key and not editable">
                (function(i,s,o,g,r,a,m){i[&apos;GoogleAnalyticsObject&apos;]=r;i[r]=i[r]||function(){
                (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
                m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
                })(window,document,&apos;script&apos;,&apos;//www.google-analytics.com/analytics.js&apos;,&apos;ga&apos;);

                ga(&apos;create&apos;, _.str.trim(&apos;<t t-esc="website.google_analytics_key"></t>&apos;), &apos;auto&apos;);
                ga(&apos;send&apos;,&apos;pageview&apos;);
            </script>
        </body>
`

	doc1 := NewDocument()
	err := doc1.ReadFromString(s)
	if err != nil {
		t.Error("etree: incorrect ReadFromString result")
	}
	doc1.GetRoot()
	fmt.Println("root", doc1.Element.Text())
	root := doc1.Root()
	fmt.Println("root", root.Tag, len(root.ChildElements()), root.ChildElements()[0].Tag)
	//s1, err := doc1.WriteToString(false)
	//if err != nil {
	//	t.Error("etree: incorrect WriteToString result")
	//}
	el := doc1.FindElement("//script")
	fmt.Println(len(el.ChildElements()), el.Text())
	el = el.Parent.ChildElements()[18]
	str, _ := el.WriteToString(false)
	//	fmt.Println("next", len(el.ChildElements()), len(el.Child),el.Child[0].)
	fmt.Println("next", str)
	//	fmt.Println("next", el.ChildElements()[0].Tag)
	el = el.GetNext()
	fmt.Println("next", len(el.ChildElements()), el.Text())
}
