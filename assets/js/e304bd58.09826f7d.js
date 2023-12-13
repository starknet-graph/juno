"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[185],{3905:(t,e,n)=>{n.d(e,{Zo:()=>u,kt:()=>c});var a=n(7294);function r(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t}function o(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(t);e&&(a=a.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,a)}return n}function l(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?o(Object(n),!0).forEach((function(e){r(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}function p(t,e){if(null==t)return{};var n,a,r=function(t,e){if(null==t)return{};var n,a,r={},o=Object.keys(t);for(a=0;a<o.length;a++)n=o[a],e.indexOf(n)>=0||(r[n]=t[n]);return r}(t,e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(t);for(a=0;a<o.length;a++)n=o[a],e.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(t,n)&&(r[n]=t[n])}return r}var i=a.createContext({}),s=function(t){var e=a.useContext(i),n=e;return t&&(n="function"==typeof t?t(e):l(l({},e),t)),n},u=function(t){var e=s(t.components);return a.createElement(i.Provider,{value:e},t.children)},m="mdxType",d={inlineCode:"code",wrapper:function(t){var e=t.children;return a.createElement(a.Fragment,{},e)}},k=a.forwardRef((function(t,e){var n=t.components,r=t.mdxType,o=t.originalType,i=t.parentName,u=p(t,["components","mdxType","originalType","parentName"]),m=s(n),k=r,c=m["".concat(i,".").concat(k)]||m[k]||d[k]||o;return n?a.createElement(c,l(l({ref:e},u),{},{components:n})):a.createElement(c,l({ref:e},u))}));function c(t,e){var n=arguments,r=e&&e.mdxType;if("string"==typeof t||r){var o=n.length,l=new Array(o);l[0]=k;var p={};for(var i in e)hasOwnProperty.call(e,i)&&(p[i]=e[i]);p.originalType=t,p[m]="string"==typeof t?t:r,l[1]=p;for(var s=2;s<o;s++)l[s]=n[s];return a.createElement.apply(null,l)}return a.createElement.apply(null,n)}k.displayName="MDXCreateElement"},4937:(t,e,n)=>{n.r(e),n.d(e,{assets:()=>i,contentTitle:()=>l,default:()=>d,frontMatter:()=>o,metadata:()=>p,toc:()=>s});var a=n(7462),r=(n(7294),n(3905));const o={slug:"/snapshots",sidebar_position:4,title:"Database Snapshots"},l=void 0,p={unversionedId:"snapshots",id:"version-0.7.0/snapshots",title:"Database Snapshots",description:"To decrease sync times, users may opt to download a Juno database snapshot.",source:"@site/versioned_docs/version-0.7.0/snapshots.md",sourceDirName:".",slug:"/snapshots",permalink:"/0.7.0/snapshots",draft:!1,tags:[],version:"0.7.0",sidebarPosition:4,frontMatter:{slug:"/snapshots",sidebar_position:4,title:"Database Snapshots"},sidebar:"tutorialSidebar",previous:{title:"Example Configuration",permalink:"/0.7.0/config"}},i={},s=[{value:"Mainnet",id:"mainnet",level:2},{value:"Goerli",id:"goerli",level:2},{value:"Goerli2",id:"goerli2",level:2},{value:"Run Juno Using Snapshot",id:"run-juno-using-snapshot",level:2}],u={toc:s},m="wrapper";function d(t){let{components:e,...n}=t;return(0,r.kt)(m,(0,a.Z)({},u,n,{components:e,mdxType:"MDXLayout"}),(0,r.kt)("p",null,"To decrease sync times, users may opt to download a Juno database snapshot.\nAfter downloading a snapshot and starting a Juno node, only recent blocks must be synced."),(0,r.kt)("h2",{id:"mainnet"},"Mainnet"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Version"),(0,r.kt)("th",{parentName:"tr",align:null},"Size"),(0,r.kt)("th",{parentName:"tr",align:null},"Block"),(0,r.kt)("th",{parentName:"tr",align:null},"Download Link"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},">=v0.6.0")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},"92 GB")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},"313975")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("a",{parentName:"td",href:"https://juno-snapshots.nethermind.dev/mainnet/juno_mainnet_v0.6.5_313975.tar"},(0,r.kt)("strong",{parentName:"a"},"juno_mainnet_313975.tar")))))),(0,r.kt)("h2",{id:"goerli"},"Goerli"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Version"),(0,r.kt)("th",{parentName:"tr",align:null},"Size"),(0,r.kt)("th",{parentName:"tr",align:null},"Block"),(0,r.kt)("th",{parentName:"tr",align:null},"Download Link"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},">=v0.6.0")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},"36 GB")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},"850192")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("a",{parentName:"td",href:"https://juno-snapshots.nethermind.dev/goerli/juno_goerli_v0.6.0_850192.tar"},(0,r.kt)("strong",{parentName:"a"},"juno_goerli_850192.tar")))))),(0,r.kt)("h2",{id:"goerli2"},"Goerli2"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Version"),(0,r.kt)("th",{parentName:"tr",align:null},"Size"),(0,r.kt)("th",{parentName:"tr",align:null},"Block"),(0,r.kt)("th",{parentName:"tr",align:null},"Download Link"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},">=v0.6.0")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},"4.6 GB")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("strong",{parentName:"td"},"139043")),(0,r.kt)("td",{parentName:"tr",align:null},(0,r.kt)("a",{parentName:"td",href:"https://juno-snapshots.nethermind.dev/goerli2/juno_goerli2_v0.6.0_139043.tar"},(0,r.kt)("strong",{parentName:"a"},"juno_goerli2_135973.tar")))))),(0,r.kt)("h2",{id:"run-juno-using-snapshot"},"Run Juno Using Snapshot"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Download Snapshot")),(0,r.kt)("p",{parentName:"li"},"Fetch a snapshot from one of the provided URLs:"),(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"wget -O juno_mainnet_313975.tar https://juno-snapshots.nethermind.dev/mainnet/juno_mainnet_v0.6.5_313975.tar\n"))),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Prepare Directory")),(0,r.kt)("p",{parentName:"li"},"Ensure you have a directory where you will store the snapshots. We will use ",(0,r.kt)("inlineCode",{parentName:"p"},"$HOME/snapshots"),"."),(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"mkdir -p $HOME/snapshots\n"))),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Extract Tarball")),(0,r.kt)("p",{parentName:"li"},"Extract the contents of the ",(0,r.kt)("inlineCode",{parentName:"p"},".tar")," file:"),(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"tar -xvf juno_mainnet_313975.tar -C $HOME/snapshots\n"))),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Run Juno")),(0,r.kt)("p",{parentName:"li"},"Execute the Docker command to run Juno, ensuring that you're using the correct snapshot path ",(0,r.kt)("inlineCode",{parentName:"p"},"$HOME/snapshots/juno_mainnet"),":"),(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"docker run -d \\\n   --name juno \\\n   -p 6060:6060 \\\n   -v $HOME/snapshots/juno_mainnet:/var/lib/juno \\\n   nethermind/juno \\\n   --http \\\n   --http-port 6060 \\\n   --http-host 0.0.0.0 \\\n   --db-path /var/lib/juno\n")))),(0,r.kt)("p",null,"After following these steps, Juno should be up and running on your machine, utilizing the provided snapshot."))}d.isMDXComponent=!0}}]);