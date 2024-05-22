"use strict";(self.webpackChunkjuno_docs=self.webpackChunkjuno_docs||[]).push([[1416],{903:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>i,default:()=>u,frontMatter:()=>s,metadata:()=>d,toc:()=>l});var o=t(4848),r=t(8453);const s={slug:"/updating_node",sidebar_position:7,title:"Updating Juno Node"},i=void 0,d={id:"updating_node",title:"Updating Juno Node",description:"Updating your Juno node is crucial to access new features, improvements, and security patches. Follow these steps to update your node to the latest version using Docker.",source:"@site/versioned_docs/version-0.11.0/updating_node.md",sourceDirName:".",slug:"/updating_node",permalink:"/updating_node",draft:!1,unlisted:!1,tags:[],version:"0.11.0",sidebarPosition:7,frontMatter:{slug:"/updating_node",sidebar_position:7,title:"Updating Juno Node"},sidebar:"main",previous:{title:"Installing on the GCP",permalink:"/installing-on-gcp"}},c={},l=[{value:"Steps to Update",id:"steps-to-update",level:3},{value:"Conclusion",id:"conclusion",level:3}];function a(e){const n={a:"a",code:"code",h3:"h3",li:"li",ol:"ol",p:"p",pre:"pre",strong:"strong",...(0,r.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.p,{children:"Updating your Juno node is crucial to access new features, improvements, and security patches. Follow these steps to update your node to the latest version using Docker."}),"\n",(0,o.jsx)(n.h3,{id:"steps-to-update",children:"Steps to Update"}),"\n",(0,o.jsxs)(n.ol,{children:["\n",(0,o.jsxs)(n.li,{children:["\n",(0,o.jsx)(n.p,{children:(0,o.jsx)(n.strong,{children:"Pull the Latest Juno Docker Image"})}),"\n",(0,o.jsxs)(n.p,{children:["First, pull the latest Juno Docker image from Nethermind's Docker repository. As an example, to update to ",(0,o.jsx)(n.code,{children:"v0.11.0-rc1"}),", use the following command:"]}),"\n"]}),"\n"]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{children:"docker pull nethermind/juno:v0.11.0-rc1\n"})}),"\n",(0,o.jsxs)(n.ol,{start:"2",children:["\n",(0,o.jsx)(n.li,{children:(0,o.jsx)(n.strong,{children:"Stop the Current Juno Container"})}),"\n"]}),"\n",(0,o.jsxs)(n.p,{children:["Stop your currently running Juno container. If you're unsure of your container's name, you can use ",(0,o.jsx)(n.code,{children:"docker ps"})," to list active containers."]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{children:"docker stop juno\n"})}),"\n",(0,o.jsxs)(n.ol,{start:"3",children:["\n",(0,o.jsx)(n.li,{children:(0,o.jsx)(n.strong,{children:"Remove the Old Container"})}),"\n"]}),"\n",(0,o.jsx)(n.p,{children:"Once the container is stopped, remove it to prevent any conflicts with the new version."}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{children:"docker rm juno\n"})}),"\n",(0,o.jsxs)(n.ol,{start:"4",children:["\n",(0,o.jsx)(n.li,{children:(0,o.jsx)(n.strong,{children:"Start a New Container with the Updated Image"})}),"\n"]}),"\n",(0,o.jsx)(n.p,{children:"Run a new container using the updated Docker image. Here's an example command, adjust it according to your setup (ports, volumes, version etc.):"}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{children:"docker run -d \\\n  --name juno \\\n  -p 6060:6060 \\\n  -v $HOME/juno:/var/lib/juno \\\n  nethermind/juno:v0.11.0-rc1 \\\n  --http \\\n  --http-port 6060 \\\n  --http-host 0.0.0.0 \\\n  --db-path /var/lib/juno \\\n  --eth-node <YOUR-ETH-NODE>\n"})}),"\n",(0,o.jsxs)(n.ol,{start:"5",children:["\n",(0,o.jsx)(n.li,{children:(0,o.jsx)(n.strong,{children:"Verify the Update"})}),"\n"]}),"\n",(0,o.jsx)(n.p,{children:"After starting the new container, verify that the node is running correctly with the updated version."}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{children:"docker logs juno\n"})}),"\n",(0,o.jsx)(n.h3,{id:"conclusion",children:"Conclusion"}),"\n",(0,o.jsxs)(n.p,{children:["You have successfully updated your Juno node to the latest version. It is now ready to be used. For more information on managing your node, visit ",(0,o.jsx)(n.a,{href:"https://github.com/NethermindEth/juno",children:"Nethermind's official GitHub repository"}),"."]})]})}function u(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(a,{...e})}):a(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>i,x:()=>d});var o=t(6540);const r={},s=o.createContext(r);function i(e){const n=o.useContext(s);return o.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function d(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:i(e.components),o.createElement(s.Provider,{value:n},e.children)}}}]);