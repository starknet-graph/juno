"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[8953],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>g});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var p=a.createContext({}),s=function(e){var t=a.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},c=function(e){var t=s(e.components);return a.createElement(p.Provider,{value:t},e.children)},u="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,p=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),u=s(n),d=r,g=u["".concat(p,".").concat(d)]||u[d]||m[d]||o;return n?a.createElement(g,i(i({ref:t},c),{},{components:n})):a.createElement(g,i({ref:t},c))}));function g(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,i=new Array(o);i[0]=d;var l={};for(var p in t)hasOwnProperty.call(t,p)&&(l[p]=t[p]);l.originalType=e,l[u]="string"==typeof e?e:r,i[1]=l;for(var s=2;s<o;s++)i[s]=n[s];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},7750:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>i,default:()=>m,frontMatter:()=>o,metadata:()=>l,toc:()=>s});var a=n(7462),r=(n(7294),n(3905));const o={slug:"/installing-on-gcp",sidebar_position:6,title:"Installing on the GCP"},i=void 0,l={unversionedId:"installing_on_gcp",id:"version-0.9.3/installing_on_gcp",title:"Installing on the GCP",description:"To install the Starknet RPC Node on the Google Cloud Platform, you can use the Juno RPC virtual machine developed by Nethermind.",source:"@site/versioned_docs/version-0.9.3/installing_on_gcp.md",sourceDirName:".",slug:"/installing-on-gcp",permalink:"/installing-on-gcp",draft:!1,tags:[],version:"0.9.3",sidebarPosition:6,frontMatter:{slug:"/installing-on-gcp",sidebar_position:6,title:"Installing on the GCP"},sidebar:"tutorialSidebar",previous:{title:"Hardware Requirements",permalink:"/hardware-requirements"}},p={},s=[{value:"Installing Starkent RPC Juno Node VM",id:"installing-starkent-rpc-juno-node-vm",level:2}],c={toc:s},u="wrapper";function m(e){let{components:t,...o}=e;return(0,r.kt)(u,(0,a.Z)({},c,o,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("p",null,"To install the Starknet RPC Node on the Google Cloud Platform, you can use the Juno RPC virtual machine developed by Nethermind."),(0,r.kt)("p",null,(0,r.kt)("strong",{parentName:"p"},"Juno")," is a golang ",(0,r.kt)("a",{parentName:"p",href:"https://starknet.io/"},"Starknet")," node implementation by ",(0,r.kt)("a",{parentName:"p",href:"https://nethermind.io/"},"Nethermind")," with the aim of decentralizing Starknet."),(0,r.kt)("h2",{id:"installing-starkent-rpc-juno-node-vm"},"Installing Starkent RPC Juno Node VM"),(0,r.kt)("p",null,"To quickly set up a Starkent RPC Juno Node VM environment on the Google Cloud Platform, follow these steps:"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Search \u201cStarknet RPC Node\u201d in ",(0,r.kt)("a",{parentName:"strong",href:"https://console.cloud.google.com/marketplace"},"Google Marketplace")," and click the LAUNCH button to start the deployment process.")),(0,r.kt)("p",{parentName:"li"},(0,r.kt)("img",{alt:"step1",src:n(3435).Z,width:"1772",height:"1506"}))),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Select the configuration for the Juno client and click the DEPLOY button.")),(0,r.kt)("p",{parentName:"li"},(0,r.kt)("img",{alt:"step2",src:n(110).Z,width:"1806",height:"1972"}))),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Post-Configuration and testing after deployment.")),(0,r.kt)("p",{parentName:"li"},(0,r.kt)("img",{alt:"step3",src:n(8645).Z,width:"1714",height:"1526"}))),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Enable Juno Auto Start During Startup")),(0,r.kt)("ol",{parentName:"li"},(0,r.kt)("li",{parentName:"ol"},"Click the newly created VM instance name to view the detail."),(0,r.kt)("li",{parentName:"ol"},"Click the Edit button."),(0,r.kt)("li",{parentName:"ol"},'Go to the "Automation" section to input the startup script as below.',(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"#! /bin/bash\nsudo /usr/local/bin/run_juno.sh\n"))),(0,r.kt)("li",{parentName:"ol"},"Click the Save button."),(0,r.kt)("li",{parentName:"ol"},"Restart the VM.")))),(0,r.kt)("ol",{start:5},(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("p",{parentName:"li"},(0,r.kt)("strong",{parentName:"p"},"Use Your Starknet RPC Juno Node")),(0,r.kt)("p",{parentName:"li"},"You can use the Juno node and access it through Rest APIs. The following is an example to verify Juno availability."),(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-bash"},'curl --location \'http://IP_Address:6060\' \\\n--header \'Content-Type: application/json\' \\\n--data \'{"jsonrpc":"2.0","method":"juno_version","params":[],"id":1}\'\n')),(0,r.kt)("p",{parentName:"li"},"The expected result is like this."),(0,r.kt)("pre",{parentName:"li"},(0,r.kt)("code",{parentName:"pre",className:"language-{"},'   "jsonrpc": "2.0",\n   "result": "v0.9.3",\n   "id": 1\n}\n')),(0,r.kt)("p",{parentName:"li"},"You can find more details from ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/NethermindEth/juno"},"https://github.com/NethermindEth/juno")))))}m.isMDXComponent=!0},3435:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/step1-c254412920cd371b4af68def7e8828a0.png"},110:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/step2-c89baac892f99fad576abe616a2ee3e2.png"},8645:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/step3-797d3db6dc7af2244bc7cc82bc1aa2b4.png"}}]);