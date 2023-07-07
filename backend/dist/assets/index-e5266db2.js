var Fe=Object.defineProperty;var Ge=(e,t,n)=>t in e?Fe(e,t,{enumerable:!0,configurable:!0,writable:!0,value:n}):e[t]=n;var ue=(e,t,n)=>(Ge(e,typeof t!="symbol"?t+"":t,n),n);(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const o of document.querySelectorAll('link[rel="modulepreload"]'))r(o);new MutationObserver(o=>{for(const i of o)if(i.type==="childList")for(const l of i.addedNodes)l.tagName==="LINK"&&l.rel==="modulepreload"&&r(l)}).observe(document,{childList:!0,subtree:!0});function n(o){const i={};return o.integrity&&(i.integrity=o.integrity),o.referrerPolicy&&(i.referrerPolicy=o.referrerPolicy),o.crossOrigin==="use-credentials"?i.credentials="include":o.crossOrigin==="anonymous"?i.credentials="omit":i.credentials="same-origin",i}function r(o){if(o.ep)return;o.ep=!0;const i=n(o);fetch(o.href,i)}})();const Qe="modulepreload",Je=function(e){return"/"+e},Ee={},J=function(t,n,r){if(!n||n.length===0)return t();const o=document.getElementsByTagName("link");return Promise.all(n.map(i=>{if(i=Je(i),i in Ee)return;Ee[i]=!0;const l=i.endsWith(".css"),s=l?'[rel="stylesheet"]':"";if(!!r)for(let m=o.length-1;m>=0;m--){const _=o[m];if(_.href===i&&(!l||_.rel==="stylesheet"))return}else if(document.querySelector(`link[href="${i}"]${s}`))return;const c=document.createElement("link");if(c.rel=l?"stylesheet":Qe,l||(c.as="script",c.crossOrigin=""),c.href=i,document.head.appendChild(c),l)return new Promise((m,_)=>{c.addEventListener("load",m),c.addEventListener("error",()=>_(new Error(`Unable to preload CSS for ${i}`)))})})).then(()=>t())};function $(){}function ge(e,t){for(const n in t)e[n]=t[n];return e}function Xe(e){return e()}function $e(){return Object.create(null)}function z(e){e.forEach(Xe)}function re(e){return typeof e=="function"}function F(e,t){return e!=e?t==t:e!==t||e&&typeof e=="object"||typeof e=="function"}let K;function Yt(e,t){return K||(K=document.createElement("a")),K.href=t,e===K.href}function Ke(e){return Object.keys(e).length===0}function Ze(e,...t){if(e==null)return $;const n=e.subscribe(...t);return n.unsubscribe?()=>n.unsubscribe():n}function Xt(e,t,n,r){if(e){const o=Ie(e,t,n,r);return e[0](o)}}function Ie(e,t,n,r){return e[1]&&r?ge(n.ctx.slice(),e[1](r(t))):n.ctx}function It(e,t,n,r){if(e[2]&&r){const o=e[2](r(n));if(t.dirty===void 0)return o;if(typeof o=="object"){const i=[],l=Math.max(t.dirty.length,o.length);for(let s=0;s<l;s+=1)i[s]=t.dirty[s]|o[s];return i}return t.dirty|o}return t.dirty}function zt(e,t,n,r,o,i){if(o){const l=Ie(t,n,r,i);e.p(l,o)}}function Bt(e){if(e.ctx.length>32){const t=[],n=e.ctx.length/32;for(let r=0;r<n;r++)t[r]=-1;return t}return-1}function te(e){return e??""}function et(e){return e&&re(e.destroy)?e.destroy:$}function S(e,t){e.appendChild(t)}function P(e,t,n){e.insertBefore(t,n||null)}function E(e){e.parentNode&&e.parentNode.removeChild(e)}function tt(e,t){for(let n=0;n<e.length;n+=1)e[n]&&e[n].d(t)}function C(e){return document.createElement(e)}function G(e){return document.createTextNode(e)}function M(){return G(" ")}function ie(){return G("")}function _e(e,t,n,r){return e.addEventListener(t,n,r),()=>e.removeEventListener(t,n,r)}function y(e,t,n){n==null?e.removeAttribute(t):e.getAttribute(t)!==n&&e.setAttribute(t,n)}function Nt(e){return e===""?null:+e}function nt(e){return Array.from(e.childNodes)}function ze(e,t){t=""+t,e.data!==t&&(e.data=t)}function Dt(e,t){e.value=t??""}function Rt(e,t,n,r){n===null?e.style.removeProperty(t):e.style.setProperty(t,n,r?"important":"")}function Tt(e,t,n){for(let r=0;r<e.options.length;r+=1){const o=e.options[r];if(o.__value===t){o.selected=!0;return}}(!n||t!==void 0)&&(e.selectedIndex=-1)}function Wt(e){const t=e.querySelector(":checked");return t&&t.__value}function Ae(e,t,n){e.classList[n?"add":"remove"](t)}function ot(e,t,{bubbles:n=!1,cancelable:r=!1}={}){const o=document.createEvent("CustomEvent");return o.initCustomEvent(e,n,r,t),o}function ne(e,t){return new e(t)}let U;function H(e){U=e}function se(){if(!U)throw new Error("Function called outside component initialization");return U}function Be(e){se().$$.on_mount.push(e)}function rt(e){se().$$.after_update.push(e)}function it(e){se().$$.on_destroy.push(e)}function st(){const e=se();return(t,n,{cancelable:r=!1}={})=>{const o=e.$$.callbacks[t];if(o){const i=ot(t,n,{cancelable:r});return o.slice().forEach(l=>{l.call(e,i)}),!i.defaultPrevented}return!0}}function xe(e,t){const n=e.$$.callbacks[t.type];n&&n.slice().forEach(r=>r.call(this,t))}const T=[],Ce=[];let W=[];const de=[],Ne=Promise.resolve();let pe=!1;function De(){pe||(pe=!0,Ne.then(Te))}function Re(){return De(),Ne}function me(e){W.push(e)}function Vt(e){de.push(e)}const fe=new Set;let N=0;function Te(){if(N!==0)return;const e=U;do{try{for(;N<T.length;){const t=T[N];N++,H(t),at(t.$$)}}catch(t){throw T.length=0,N=0,t}for(H(null),T.length=0,N=0;Ce.length;)Ce.pop()();for(let t=0;t<W.length;t+=1){const n=W[t];fe.has(n)||(fe.add(n),n())}W.length=0}while(T.length);for(;de.length;)de.pop()();pe=!1,fe.clear(),H(e)}function at(e){if(e.fragment!==null){e.update(),z(e.before_update);const t=e.dirty;e.dirty=[-1],e.fragment&&e.fragment.p(e.ctx,t),e.after_update.forEach(me)}}function lt(e){const t=[],n=[];W.forEach(r=>e.indexOf(r)===-1?t.push(r):n.push(r)),n.forEach(r=>r()),W=t}const ee=new Set;let I;function be(){I={r:0,c:[],p:I}}function ye(){I.r||z(I.c),I=I.p}function k(e,t){e&&e.i&&(ee.delete(e),e.i(t))}function q(e,t,n,r){if(e&&e.o){if(ee.has(e))return;ee.add(e),I.c.push(()=>{ee.delete(e),r&&(n&&e.d(1),r())}),e.o(t)}else r&&r()}const Ht=typeof window<"u"?window:typeof globalThis<"u"?globalThis:global;function We(e,t){const n={},r={},o={$$scope:1};let i=e.length;for(;i--;){const l=e[i],s=t[i];if(s){for(const a in l)a in s||(r[a]=1);for(const a in s)o[a]||(n[a]=s[a],o[a]=1);e[i]=s}else for(const a in l)o[a]=1}for(const l in r)l in n||(n[l]=void 0);return n}function Ve(e){return typeof e=="object"&&e!==null?e:{}}function Ut(e,t,n){const r=e.$$.props[t];r!==void 0&&(e.$$.bound[r]=n,n(e.$$.ctx[r]))}function j(e){e&&e.c()}function A(e,t,n,r){const{fragment:o,after_update:i}=e.$$;o&&o.m(t,n),r||me(()=>{const l=e.$$.on_mount.map(Xe).filter(re);e.$$.on_destroy?e.$$.on_destroy.push(...l):z(l),e.$$.on_mount=[]}),i.forEach(me)}function x(e,t){const n=e.$$;n.fragment!==null&&(lt(n.after_update),z(n.on_destroy),n.fragment&&n.fragment.d(t),n.on_destroy=n.fragment=null,n.ctx=[])}function ct(e,t){e.$$.dirty[0]===-1&&(T.push(e),De(),e.$$.dirty.fill(0)),e.$$.dirty[t/31|0]|=1<<t%31}function ae(e,t,n,r,o,i,l,s=[-1]){const a=U;H(e);const c=e.$$={fragment:null,ctx:[],props:i,update:$,not_equal:o,bound:$e(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(t.context||(a?a.$$.context:[])),callbacks:$e(),dirty:s,skip_bound:!1,root:t.target||a.$$.root};l&&l(c.root);let m=!1;if(c.ctx=n?n(e,t.props||{},(_,f,...g)=>{const v=g.length?g[0]:f;return c.ctx&&o(c.ctx[_],c.ctx[_]=v)&&(!c.skip_bound&&c.bound[_]&&c.bound[_](v),m&&ct(e,_)),f}):[],c.update(),m=!0,z(c.before_update),c.fragment=r?r(c.ctx):!1,t.target){if(t.hydrate){const _=nt(t.target);c.fragment&&c.fragment.l(_),_.forEach(E)}else c.fragment&&c.fragment.c();t.intro&&k(e.$$.fragment),A(e,t.target,t.anchor,t.customElement),Te()}H(a)}class le{$destroy(){x(this,1),this.$destroy=$}$on(t,n){if(!re(n))return $;const r=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return r.push(n),()=>{const o=r.indexOf(n);o!==-1&&r.splice(o,1)}}$set(t){this.$$set&&!Ke(t)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}function Z(e){if(!e)throw Error("Parameter args is required");if(!e.component==!e.asyncComponent)throw Error("One and only one of component and asyncComponent is required");if(e.component&&(e.asyncComponent=()=>Promise.resolve(e.component)),typeof e.asyncComponent!="function")throw Error("Parameter asyncComponent must be a function");if(e.conditions){Array.isArray(e.conditions)||(e.conditions=[e.conditions]);for(let n=0;n<e.conditions.length;n++)if(!e.conditions[n]||typeof e.conditions[n]!="function")throw Error("Invalid parameter conditions["+n+"]")}return e.loadingComponent&&(e.asyncComponent.loading=e.loadingComponent,e.asyncComponent.loadingParams=e.loadingParams||void 0),{component:e.asyncComponent,userData:e.userData,conditions:e.conditions&&e.conditions.length?e.conditions:void 0,props:e.props&&Object.keys(e.props).length?e.props:{},_sveltesparouter:!0}}const D=[];function He(e,t){return{subscribe:ce(e,t).subscribe}}function ce(e,t=$){let n;const r=new Set;function o(s){if(F(e,s)&&(e=s,n)){const a=!D.length;for(const c of r)c[1](),D.push(c,e);if(a){for(let c=0;c<D.length;c+=2)D[c][0](D[c+1]);D.length=0}}}function i(s){o(s(e))}function l(s,a=$){const c=[s,a];return r.add(c),r.size===1&&(n=t(o)||$),s(e),()=>{r.delete(c),r.size===0&&n&&(n(),n=null)}}return{set:o,update:i,subscribe:l}}function Ue(e,t,n){const r=!Array.isArray(e),o=r?[e]:e,i=t.length<2;return He(n,l=>{let s=!1;const a=[];let c=0,m=$;const _=()=>{if(c)return;m();const g=t(r?a[0]:a,l);i?l(g):m=re(g)?g:$},f=o.map((g,v)=>Ze(g,O=>{a[v]=O,c&=~(1<<v),s&&_()},()=>{c|=1<<v}));return s=!0,_(),function(){z(f),m(),s=!1}})}function ut(e,t){if(e instanceof RegExp)return{keys:!1,pattern:e};var n,r,o,i,l=[],s="",a=e.split("/");for(a[0]||a.shift();o=a.shift();)n=o[0],n==="*"?(l.push("wild"),s+="/(.*)"):n===":"?(r=o.indexOf("?",1),i=o.indexOf(".",1),l.push(o.substring(1,~r?r:~i?i:o.length)),s+=~r&&!~i?"(?:/([^/]+?))?":"/([^/]+?)",~i&&(s+=(~r?"?":"")+"\\"+o.substring(i))):s+="/"+o;return{keys:l,pattern:new RegExp("^"+s+(t?"(?=$|/)":"/?$"),"i")}}function ft(e){let t,n,r;const o=[e[2]];var i=e[0];function l(s){let a={};for(let c=0;c<o.length;c+=1)a=ge(a,o[c]);return{props:a}}return i&&(t=ne(i,l()),t.$on("routeEvent",e[7])),{c(){t&&j(t.$$.fragment),n=ie()},m(s,a){t&&A(t,s,a),P(s,n,a),r=!0},p(s,a){const c=a&4?We(o,[Ve(s[2])]):{};if(a&1&&i!==(i=s[0])){if(t){be();const m=t;q(m.$$.fragment,1,0,()=>{x(m,1)}),ye()}i?(t=ne(i,l()),t.$on("routeEvent",s[7]),j(t.$$.fragment),k(t.$$.fragment,1),A(t,n.parentNode,n)):t=null}else i&&t.$set(c)},i(s){r||(t&&k(t.$$.fragment,s),r=!0)},o(s){t&&q(t.$$.fragment,s),r=!1},d(s){s&&E(n),t&&x(t,s)}}}function dt(e){let t,n,r;const o=[{params:e[1]},e[2]];var i=e[0];function l(s){let a={};for(let c=0;c<o.length;c+=1)a=ge(a,o[c]);return{props:a}}return i&&(t=ne(i,l()),t.$on("routeEvent",e[6])),{c(){t&&j(t.$$.fragment),n=ie()},m(s,a){t&&A(t,s,a),P(s,n,a),r=!0},p(s,a){const c=a&6?We(o,[a&2&&{params:s[1]},a&4&&Ve(s[2])]):{};if(a&1&&i!==(i=s[0])){if(t){be();const m=t;q(m.$$.fragment,1,0,()=>{x(m,1)}),ye()}i?(t=ne(i,l()),t.$on("routeEvent",s[6]),j(t.$$.fragment),k(t.$$.fragment,1),A(t,n.parentNode,n)):t=null}else i&&t.$set(c)},i(s){r||(t&&k(t.$$.fragment,s),r=!0)},o(s){t&&q(t.$$.fragment,s),r=!1},d(s){s&&E(n),t&&x(t,s)}}}function pt(e){let t,n,r,o;const i=[dt,ft],l=[];function s(a,c){return a[1]?0:1}return t=s(e),n=l[t]=i[t](e),{c(){n.c(),r=ie()},m(a,c){l[t].m(a,c),P(a,r,c),o=!0},p(a,[c]){let m=t;t=s(a),t===m?l[t].p(a,c):(be(),q(l[m],1,1,()=>{l[m]=null}),ye(),n=l[t],n?n.p(a,c):(n=l[t]=i[t](a),n.c()),k(n,1),n.m(r.parentNode,r))},i(a){o||(k(n),o=!0)},o(a){q(n),o=!1},d(a){l[t].d(a),a&&E(r)}}}function Pe(){const e=window.location.href.indexOf("#/");let t=e>-1?window.location.href.substr(e+1):"/";const n=t.indexOf("?");let r="";return n>-1&&(r=t.substr(n+1),t=t.substr(0,n)),{location:t,querystring:r}}const ve=He(null,function(t){t(Pe());const n=()=>{t(Pe())};return window.addEventListener("hashchange",n,!1),function(){window.removeEventListener("hashchange",n,!1)}});Ue(ve,e=>e.location);Ue(ve,e=>e.querystring);const je=ce(void 0);async function Ft(e){if(!e||e.length<1||e.charAt(0)!="/"&&e.indexOf("#/")!==0)throw Error("Invalid parameter location");await Re(),history.replaceState({...history.state,__svelte_spa_router_scrollX:window.scrollX,__svelte_spa_router_scrollY:window.scrollY},void 0),window.location.hash=(e.charAt(0)=="#"?"":"#")+e}function mt(e,t){if(t=Le(t),!e||!e.tagName||e.tagName.toLowerCase()!="a")throw Error('Action "link" can only be used with <a> tags');return Oe(e,t),{update(n){n=Le(n),Oe(e,n)}}}function ht(e){e?window.scrollTo(e.__svelte_spa_router_scrollX,e.__svelte_spa_router_scrollY):window.scrollTo(0,0)}function Oe(e,t){let n=t.href||e.getAttribute("href");if(n&&n.charAt(0)=="/")n="#"+n;else if(!n||n.length<2||n.slice(0,2)!="#/")throw Error('Invalid value for "href" attribute: '+n);e.setAttribute("href",n),e.addEventListener("click",r=>{r.preventDefault(),t.disabled||gt(r.currentTarget.getAttribute("href"))})}function Le(e){return e&&typeof e=="string"?{href:e}:e||{}}function gt(e){history.replaceState({...history.state,__svelte_spa_router_scrollX:window.scrollX,__svelte_spa_router_scrollY:window.scrollY},void 0),window.location.hash=e}function _t(e,t,n){let{routes:r={}}=t,{prefix:o=""}=t,{restoreScrollState:i=!1}=t;class l{constructor(p,d){if(!d||typeof d!="function"&&(typeof d!="object"||d._sveltesparouter!==!0))throw Error("Invalid component object");if(!p||typeof p=="string"&&(p.length<1||p.charAt(0)!="/"&&p.charAt(0)!="*")||typeof p=="object"&&!(p instanceof RegExp))throw Error('Invalid value for "path" argument - strings must start with / or *');const{pattern:w,keys:u}=ut(p);this.path=p,typeof d=="object"&&d._sveltesparouter===!0?(this.component=d.component,this.conditions=d.conditions||[],this.userData=d.userData,this.props=d.props||{}):(this.component=()=>Promise.resolve(d),this.conditions=[],this.props={}),this._pattern=w,this._keys=u}match(p){if(o){if(typeof o=="string")if(p.startsWith(o))p=p.substr(o.length)||"/";else return null;else if(o instanceof RegExp){const b=p.match(o);if(b&&b[0])p=p.substr(b[0].length)||"/";else return null}}const d=this._pattern.exec(p);if(d===null)return null;if(this._keys===!1)return d;const w={};let u=0;for(;u<this._keys.length;){try{w[this._keys[u]]=decodeURIComponent(d[u+1]||"")||null}catch{w[this._keys[u]]=null}u++}return w}async checkConditions(p){for(let d=0;d<this.conditions.length;d++)if(!await this.conditions[d](p))return!1;return!0}}const s=[];r instanceof Map?r.forEach((h,p)=>{s.push(new l(p,h))}):Object.keys(r).forEach(h=>{s.push(new l(h,r[h]))});let a=null,c=null,m={};const _=st();async function f(h,p){await Re(),_(h,p)}let g=null,v=null;i&&(v=h=>{h.state&&(h.state.__svelte_spa_router_scrollY||h.state.__svelte_spa_router_scrollX)?g=h.state:g=null},window.addEventListener("popstate",v),rt(()=>{ht(g)}));let O=null,L=null;const B=ve.subscribe(async h=>{O=h;let p=0;for(;p<s.length;){const d=s[p].match(h.location);if(!d){p++;continue}const w={route:s[p].path,location:h.location,querystring:h.querystring,userData:s[p].userData,params:d&&typeof d=="object"&&Object.keys(d).length?d:null};if(!await s[p].checkConditions(w)){n(0,a=null),L=null,f("conditionsFailed",w);return}f("routeLoading",Object.assign({},w));const u=s[p].component;if(L!=u){u.loading?(n(0,a=u.loading),L=u,n(1,c=u.loadingParams),n(2,m={}),f("routeLoaded",Object.assign({},w,{component:a,name:a.name,params:c}))):(n(0,a=null),L=null);const b=await u();if(h!=O)return;n(0,a=b&&b.default||b),L=u}d&&typeof d=="object"&&Object.keys(d).length?n(1,c=d):n(1,c=null),n(2,m=s[p].props),f("routeLoaded",Object.assign({},w,{component:a,name:a.name,params:c})).then(()=>{je.set(c)});return}n(0,a=null),L=null,je.set(void 0)});it(()=>{B(),v&&window.removeEventListener("popstate",v)});function Y(h){xe.call(this,e,h)}function Q(h){xe.call(this,e,h)}return e.$$set=h=>{"routes"in h&&n(3,r=h.routes),"prefix"in h&&n(4,o=h.prefix),"restoreScrollState"in h&&n(5,i=h.restoreScrollState)},e.$$.update=()=>{e.$$.dirty&32&&(history.scrollRestoration=i?"manual":"auto")},[a,c,m,r,o,i,Y,Q]}class bt extends le{constructor(t){super(),ae(this,t,_t,pt,F,{routes:3,prefix:4,restoreScrollState:5})}}var V=(e=>(e.about="about",e.giveMeASine="giveMeASine",e.home="home",e.magicSquare="magicSquare",e.none="none",e))(V||{});function yt(e){switch(e){case"about":return"about";case"giveMeASine":return"giveMeASine";case"home":return"home";case"magicSquare":return"magicSquare";default:return"none"}}function Gt(e){switch(e){case"about":return"/about";case"home":return"/";case"giveMeASine":return"/give_me_a_sine";case"magicSquare":return"/magic_square";case"none":default:return"/"}}const oe=ce("home");function vt(e){let t,n,r,o,i,l;return{c(){t=C("a"),n=C("button"),r=G(e[3]),y(n,"class","link_button svelte-yriz3c"),y(t,"href",e[1]),y(t,"target","_blank"),y(t,"class",o=te(`link ${e[2]}`)+" svelte-yriz3c")},m(s,a){P(s,t,a),S(t,n),S(n,r),i||(l=_e(t,"click",e[7]),i=!0)},p(s,a){a&8&&ze(r,s[3]),a&2&&y(t,"href",s[1]),a&4&&o!==(o=te(`link ${s[2]}`)+" svelte-yriz3c")&&y(t,"class",o)},d(s){s&&E(t),i=!1,l()}}}function wt(e){let t,n,r,o,i,l;return{c(){t=C("a"),n=C("button"),r=G(e[3]),y(n,"class","link_button svelte-yriz3c"),y(t,"href",e[1]),y(t,"class",o=te(`link ${e[2]}`)+" svelte-yriz3c")},m(s,a){P(s,t,a),S(t,n),S(n,r),i||(l=[et(mt.call(null,t)),_e(t,"click",e[6])],i=!0)},p(s,a){a&8&&ze(r,s[3]),a&2&&y(t,"href",s[1]),a&4&&o!==(o=te(`link ${s[2]}`)+" svelte-yriz3c")&&y(t,"class",o)},d(s){s&&E(t),i=!1,z(l)}}}function St(e){let t;function n(i,l){return i[0]?wt:vt}let r=n(e),o=r(e);return{c(){o.c(),t=ie()},m(i,l){o.m(i,l),P(i,t,l)},p(i,[l]){r===(r=n(i))&&o?o.p(i,l):(o.d(1),o=r(i),o&&(o.c(),o.m(t.parentNode,t)))},i:$,o:$,d(i){o.d(i),i&&E(t)}}}function kt(e,t,n){oe.subscribe(f=>f);function r(f){oe.update(g=>f)}function o(f){var g=V.home;switch(f){case"/about":g=V.about;break;case"/magic_square":g=V.magicSquare;break;case"/give_me_a_sine":g=V.giveMeASine;break}localStorage.setItem("ns_site_section",g),r(g)}let{sameOrigin:i=!0}=t,{href:l="/"}=t,{className:s=""}=t,{title:a="Home"}=t,{onClick:c=f=>{f.stopPropagation()}}=t;const m=f=>{o(l),c(f)},_=f=>{c(f)};return e.$$set=f=>{"sameOrigin"in f&&n(0,i=f.sameOrigin),"href"in f&&n(1,l=f.href),"className"in f&&n(2,s=f.className),"title"in f&&n(3,a=f.title),"onClick"in f&&n(4,c=f.onClick)},[i,l,s,a,c,o,m,_]}class R extends le{constructor(t){super(),ae(this,t,kt,St,F,{sameOrigin:0,href:1,className:2,title:3,onClick:4})}}const qt={about:{personalProejects:"Personal projects",technicalExperience:"Technical experience",openWasm:"Open the WebAssembly build"},app:{nav:{home:"Home",about:"About",magicSquare:"Magic square",giveMeASine:"Give me a sine",none:""}},gmas:{aboveColor:"above color",belowColor:"below color",graphColor:"graph color",height:"height",width:"width"},home:{title:"Nate's New Website",about:"About",about_1:"An overview of what I do code-wise",about_2:'Please enjoy these AI-generated images of me in "natural light"',magicSquare:"Magic square",magicSquare_1:"Interactive 3D with",magicSquare_2:"Exploring concepts from",magicSquare_3:"modular synthesis",giveMeASine:"Give me a sine",giveMeASine_1:"All graphs copy-and-paste-able",giveMeASine_2:"practice for",intro:{1:"I like to learn.",2:"I wanted to learn how to combine",3:"using",4:"So I built this new website."},whatsHere:"What's here?"},magicSquare:{color:{animation:"animation",in:"in",out:"out",fix:"fix",speed:"speed"},controlModule:{empty:"empty"},controlRack:{color:"color",drawPattern:"pattern",geometry:"geometry",lfo:"lfo",modules:"modules",presets:"presets",radius:"radius",rotation:"rotation",translation:"translation"},drawPattern:{animation:"animation",direction:"direction",count:"count",in:"in",out:"out",fix:"fix",offset:"offset",order:"order",rotate:"rotate",speed:"speed",translate:"translate"},geometry:{shape:"shape",radius:"radius",range:"range",ngon:"ngon",star:"star",coolS:"cool s",tetrahedron:"tetrahedron",cube:"cube",octahedron:"octahedron",dodecahedron:"dodecahedron",icosahedron:"icosahedron"},main:{active:"active",amplitude:"amplitude",base:"base",count:"count",frequency:"speed",offset:"offset",phase:"phase",speed:"speed",spread:"spread",step:"step"},presets:{bank:"bank",curr:"current",load:"load",next:"next",preset:"preset",save:"save"},lfo:{destination:"destination",minimum:"minimum",pitchBase:"pitch base",pitchSpread:"pitch spread",pitchX:"pitch X",pitchY:"pitch Y",rollBase:"roll base",rollSpread:"roll spread",rollX:"roll X",rollY:"roll Y",yawBase:"yaw base",yawSpread:"yaw spread",yawX:"yaw X",yawY:"yaw Y",radius:"radius",radiusBase:"radius base",radiusStep:"radius step",rotation:"rotation",shape:"shape",translation:"translation",xBase:"X base",xSpread:"X spread",yBase:"Y base",ySpread:"Y spread"},mouseTracking:{invert:"invert",mouse:"mouse",none:"none",on:"on",off:"off"},rotation:{pitch:"pitch",roll:"roll",yaw:"yaw"},select:{left:"left",right:"right",color:"color",drawPattern:"pattern",geometry:"geometry",lfo:"lfo",presets:"presets",radius:"radius",rotation:"rotation",translation:"translation"},warning:{title:"Epilepsy Warning",body_1:"Please be aware that this software contains sequences of flashing lights which may trigger seizures for people with photosensitive epilepsy. Viewer discretion is advised.",body_2:"If you or anyone in your household has an epileptic condition, please consult a medical professional before using this software. If you experience dizziness, altered vision, eye or muscle twitches, loss of awareness, disorientation, or any involuntary movement or convulsion while using the software, immediately discontinue use and consult a medical professional.",body_3:"Please ensure you enjoy Magic Square in a well-lit environment and take frequent breaks. Your safety is important to us.",go_home:"Go To Home",accept_and_continue:"Accept & Continue"}}},Et={about:{personalProejects:"Proyectos personales",technicalExperience:"Experiencia técnica",openWasm:"Abrir el build WebAssembly"},app:{nav:{home:"Inicio",about:"Sobre mi",magicSquare:"Cuadro mágico",giveMeASine:"Dame uno seno",none:""}},gmas:{aboveColor:"color ariba",belowColor:"color abajo",graphColor:"color grafico",height:"altura",width:"ancho"},home:{title:"El sitio nuevo de Nate",about:"Sobre mi",about_1:"Lo que hago por lo general con codigo",about_2:'Por favor, disfrute de estas imágenes generadas por inteligencia artificial de mí en "luz natural"',magicSquare:"Cuadro mágico",magicSquare_1:"3D interactivo con",magicSquare_2:"Explorando conceptos de",magicSquare_3:"síntesis modular",giveMeASine:"Dame uno seno",giveMeASine_1:"Todos gráficos se pueden copiar y pegar",giveMeASine_2:"preparación para",intro:{1:"Me gusta apprender.",2:"Queria apprender como mezclar",3:"usando",4:"Entonces, hizo ese sitio nuevo."},whatsHere:"Que hay aqui?"},magicSquare:{color:{animation:"animación",in:"int",out:"ext",fix:"fijo",speed:"velocidad"},controlModule:{empty:"vacío"},controlRack:{color:"color",drawPattern:"dispersión",geometry:"geometría",lfo:"lfo",modules:"módulos",presets:"preajustes",radius:"radio",rotation:"rotación",translation:"traslado"},drawPattern:{animation:"animación",direction:"dirección",order:"orden",in:"int",offset:"offset",out:"ext",fix:"fijo",rotate:"girar",translate:"traslado"},geometry:{shape:"forma",radius:"radio",range:"serie",ngon:"ngon",star:"estrella",coolS:"s cool",tetrahedron:"tetraedro",cube:"cubo",octahedron:"octaedro",dodecahedron:"dodecaedro",icosahedron:"icosaedro"},lfo:{destination:"destino",minimum:"mínimo",pitchBase:"alabeo base",pitchSpread:"alabeo escalón",pitchX:"alabeo X",pitchY:"alabeo Y",rollBase:"balanceo base",rollSpread:"balanceo escalón",rollX:"balanceo X",rollY:"balanceo Y",yawBase:"guiñada base",yawSpread:"guiñada escalón",yawX:"guiñada X",yawY:"guiñada Y",radius:"radio",radiusBase:"base de radio",radiusStep:"escalón de radio",rotation:"rotación",shape:"forma",translation:"traslado",xBase:"X base",xSpread:"X escalón",yBase:"Y base",ySpread:"Y escalón"},main:{active:"activo",amplitude:"amplitud",base:"base",count:"cuenta",frequency:"velocidad",minimum:"mínimo",offset:"offset",phase:"fase",speed:"velocidad",spread:"escalón",step:"paso"},mouseTracking:{invert:"invertir",mouse:"ratón",none:"ninguno",on:"encendido",off:"apagado"},presets:{bank:"banco",curr:"actual",load:"carga",next:"proxima",preset:"preajuste",save:"guarda"},rotation:{pitch:"alabeo",roll:"balanceo",yaw:"guiñada"},select:{left:"izquierda",right:"derecha",color:"color",drawPattern:"dispersión",geometry:"geometría",lfo:"lfo",presets:"preajustes",radius:"radio",rotation:"rotación",translation:"traslado"},warning:{title:"Advertencia de Epilepsia ",body_1:"Por favor, tenga en cuenta que este software contiene secuencias de luces intermitentes que pueden provocar convulsiones a personas con epilepsia fotosensitiva. Se recomienda discreción al espectador.",body_2:"Si usted o alguien en su hogar padece una condición epiléptica, consulte a un profesional médico antes de usar este software. Si experimenta mareos, visión alterada, parpadeo ocular o muscular, pérdida de conciencia, desorientación, o cualquier movimiento o convulsión involuntaria mientras usa el software, deje de usarlo inmediatamente y consulte a un profesional médico.",body_3:"Asegúrese de usar en un ambiente bien iluminado y tome descansos frecuentes. Su seguridad es importante para nosotros.",go_home:"Regresar a inicio",accept_and_continue:"Aceptar y continuar"}}},$t={about:{personalProejects:"Projets personnels",technicalExperience:"Expérience technique",openWasm:"Ouvrir le build WebAssembly"},app:{nav:{home:"Accueil",about:"De moi",magicSquare:"Carré magique",giveMeASine:"Donne-moi un sinus",none:""}},gmas:{aboveColor:"couleur dessus",belowColor:"couleur dessous",graphColor:"couleur graph",height:"hauteur",width:"largeur"},home:{title:"Le site nouveau de Nate",about:"De moi",about_1:"Ce que je fait con code",about_2:`S'il vous plaît, appréciez ces images de moi en "lumière naturelle" générées par intelligence artificielle`,magicSquare:"Carré magique",magicSquare_1:"3D interactive avec",magicSquare_2:"Exploration de concepts de",magicSquare_3:"synthèse modulaire",giveMeASine:"Donne-moi un sinus",giveMeASine_1:"On peut copier et coller tous les graphiques",giveMeASine_2:"préparation pour",intro:{1:"J'aime apprendre.",2:"Je voulais apprendre comment mélanger",3:"en utilisant",4:"Alors, j'ai fait ce site nouveau."},whatsHere:"Qu'est-ce qu'il y a ici?"},magicSquare:{color:{animation:"animation",in:"int",out:"ext",fix:"fixé",speed:"vitesse"},controlModule:{empty:"vide"},controlRack:{color:"couleur",drawPattern:"dispersion",geometry:"géométrie",lfo:"lfo",modules:"modules",presets:"préréglages",radius:"rayon",rotation:"rotation",translation:"translation"},drawPattern:{animation:"animation",direction:"direction",order:"ordre",in:"int",offset:"offset",out:"ext",fix:"fixé",rotate:"tourner",translate:"traduir"},geometry:{shape:"forme",radius:"rayon",range:"gamme",ngon:"ngon",star:"étoile",coolS:"s cool",tetrahedron:"tétraèdre",cube:"cube",octahedron:"octaèdre",dodecahedron:"dodécaèdre",icosahedron:"icosaèdre"},lfo:{destination:"destination",minimum:"mínimo",pitchBase:"tangage base",pitchSpread:"pas de tangage",pitchX:"tangage X",pitchY:"tangage Y",rollBase:"roulis base",rollSpread:"pas de roulis",rollX:"roulis X",rollY:"roulis Y",yawBase:"lacet base",yawSpread:"pas de lacet",yawX:"lacet X",yawY:"lacet Y",radius:"rayon",radiusBase:"rayon base",radiusStep:"pas de rayon",rotation:"rotation",shape:"forme",translation:"translation",xBase:"X base",xSpread:"pas de X",yBase:"Y base",ySpread:"pas de Y"},main:{active:"actif",amplitude:"amplitude",base:"base",count:"cuenta",frequency:"vitesse",minimum:"minimum",offset:"offset",phase:"phase",speed:"vitesse",spread:"pas de",step:"pas"},mouseTracking:{invert:"inverse",mouse:"souris",none:"aucun",on:"marche",off:"arrêt"},presets:{bank:"banque",curr:"actuel",load:"charge",next:"prochain",preset:"préréglage",save:"sauve"},rotation:{pitch:"tangage",roll:"roulis",yaw:"lacet"},select:{left:"gauche",right:"droite",color:"couleur",drawPattern:"dispersion",geometry:"géométrie",lfo:"lfo",presets:"préréglages",radius:"rayon",rotation:"rotation",translation:"translation"},warning:{title:"Avertissement sur l'Épilepsie",body_1:"Veuillez noter que ce logiciel contient des séquences de lumières clignotantes qui peuvent déclencher des crises d'épilepsie chez les personnes atteintes d'épilepsie photosensible. La prudence du spectateur est conseillée.",body_2:"Si vous ou quelqu'un dans votre foyer souffrez d'une condition épileptique, veuillez consulter un professionnel de la santé avant d'utiliser ce logiciel. Si vous ressentez des étourdissements, une vision altérée, des tics oculaires ou musculaires, une perte de conscience, une désorientation, ou tout mouvement ou convulsion involontaire lors de l'utilisation du logiciel, arrêtez immédiatement son utilisation et consultez un professionnel de la santé.",body_3:"Assurez-vous de utiliser dans un environnement bien éclairé et de faire des pauses fréquentes. Votre sécurité est importante pour nous.",go_home:"Retour à l'accueil",accept_and_continue:"Accepter et continuer"}},misc:{}};var X=(e=>(e.en="en",e.es="es",e.fr="fr",e))(X||{});class At{constructor(t){ue(this,"prefix");ue(this,"locales",{en:qt,es:Et,fr:$t});this.prefix=t.split("/")}t(t,n){const r=this.prefix.concat(t.split("/"));var i=(this.locales[n]?this.locales[n]:this.locales.en)[r.shift()];if(r.length&&r.forEach(l=>{i=(i||{})[l]}),typeof i=="string")return i;switch(n){case"en":return"Translation not found";case"es":return"Traducción no encontrada";case"fr":return"Traduction non trouvée"}}}const he=ce(null);function Me(e,t,n){const r=e.slice();return r[3]=t[n],r}function Ye(e){let t,n=e[3]+"",r,o,i,l;function s(){return e[2](e[3])}return{c(){t=C("button"),r=G(n),o=M(),y(t,"class","lang_select_opt mt-0 svelte-a5vudu"),Ae(t,"selected",X[e[3]]===e[0])},m(a,c){P(a,t,c),S(t,r),S(t,o),i||(l=_e(t,"click",s),i=!0)},p(a,c){e=a,c&1&&Ae(t,"selected",X[e[3]]===e[0])},d(a){a&&E(t),i=!1,l()}}}function xt(e){let t,n,r=Object.keys(X),o=[];for(let i=0;i<r.length;i+=1)o[i]=Ye(Me(e,r,i));return{c(){t=C("section"),n=C("div");for(let i=0;i<o.length;i+=1)o[i].c();y(n,"class","lang_select grow")},m(i,l){P(i,t,l),S(t,n);for(let s=0;s<o.length;s+=1)o[s]&&o[s].m(n,null)},p(i,[l]){if(l&3){r=Object.keys(X);let s;for(s=0;s<r.length;s+=1){const a=Me(i,r,s);o[s]?o[s].p(a,l):(o[s]=Ye(a),o[s].c(),o[s].m(n,null))}for(;s<o.length;s+=1)o[s].d(1);o.length=r.length}},i:$,o:$,d(i){i&&E(t),tt(o,i)}}}function Ct(e,t,n){let r;he.subscribe(l=>n(0,r=l));function o(l){localStorage.setItem("lang",X[l]),he.update(s=>X[l])}return Be(()=>{const l=localStorage.getItem("lang");o(typeof l=="string"?l:X.en)}),[r,o,l=>o(l)]}class Pt extends le{constructor(t){super(),ae(this,t,Ct,xt,F,{})}}function jt(e){let t,n,r,o,i,l,s,a,c,m,_,f,g,v,O,L,B,Y,Q,h,p,d,w;return n=new R({props:{href:"/",title:e[1].t("nav/home",e[0])}}),o=new R({props:{href:"/about",title:e[1].t("nav/about",e[0])}}),l=new R({props:{href:"/magic_square",title:e[1].t("nav/magicSquare",e[0])}}),a=new R({props:{href:"/give_me_a_sine",title:e[1].t("nav/giveMeASine",e[0])}}),_=new bt({props:{routes:e[2]}}),O=new R({props:{href:"https://github.com/nathanielBellamy",title:"github.com/nathanielBellamy",sameOrigin:!1}}),Y=new R({props:{href:"mailto:nbschieber@gmail.com",title:"nbschieber@gmail.com",sameOrigin:!1}}),d=new Pt({}),{c(){t=C("nav"),j(n.$$.fragment),r=M(),j(o.$$.fragment),i=M(),j(l.$$.fragment),s=M(),j(a.$$.fragment),c=M(),m=C("main"),j(_.$$.fragment),f=M(),g=C("footer"),v=C("div"),j(O.$$.fragment),L=M(),B=C("div"),j(Y.$$.fragment),Q=M(),h=C("div"),h.textContent="PORTLAND, OR",p=M(),j(d.$$.fragment),y(t,"class","nav_bar flex items-center gap-2 svelte-ll8jqi"),y(m,"class","rounded-md flex flex-col justify-start pb-20 md:pb-0"),y(v,"class","grow"),y(B,"class","grow"),y(h,"class","city svelte-ll8jqi"),y(g,"class","flex flex-col space-between items-stretch pt-2 pb-2 md:flex-row md:pb-0")},m(u,b){P(u,t,b),A(n,t,null),S(t,r),A(o,t,null),S(t,i),A(l,t,null),S(t,s),A(a,t,null),P(u,c,b),P(u,m,b),A(_,m,null),P(u,f,b),P(u,g,b),S(g,v),A(O,v,null),S(g,L),S(g,B),A(Y,B,null),S(g,Q),S(g,h),S(g,p),A(d,g,null),w=!0},p(u,[b]){const we={};b&1&&(we.title=u[1].t("nav/home",u[0])),n.$set(we);const Se={};b&1&&(Se.title=u[1].t("nav/about",u[0])),o.$set(Se);const ke={};b&1&&(ke.title=u[1].t("nav/magicSquare",u[0])),l.$set(ke);const qe={};b&1&&(qe.title=u[1].t("nav/giveMeASine",u[0])),a.$set(qe)},i(u){w||(k(n.$$.fragment,u),k(o.$$.fragment,u),k(l.$$.fragment,u),k(a.$$.fragment,u),k(_.$$.fragment,u),k(O.$$.fragment,u),k(Y.$$.fragment,u),k(d.$$.fragment,u),w=!0)},o(u){q(n.$$.fragment,u),q(o.$$.fragment,u),q(l.$$.fragment,u),q(a.$$.fragment,u),q(_.$$.fragment,u),q(O.$$.fragment,u),q(Y.$$.fragment,u),q(d.$$.fragment,u),w=!1},d(u){u&&E(t),x(n),x(o),x(l),x(a),u&&E(c),u&&E(m),x(_),u&&E(f),u&&E(g),x(O),x(Y),x(d)}}}function Ot(e,t,n){oe.subscribe(l=>l);let r=new At("app"),o;he.subscribe(l=>n(0,o=l));const i={"/":Z({asyncComponent:()=>J(()=>import("./Home-bb7a1c56.js"),["assets/Home-bb7a1c56.js","assets/Home-f88b5e0f.css"])}),"/about":Z({asyncComponent:()=>J(()=>import("./About-df2412f5.js"),["assets/About-df2412f5.js","assets/About-0f61a0ab.css"])}),"/give_me_a_sine":Z({asyncComponent:()=>J(()=>import("./GiveMeASine-6c993962.js"),["assets/GiveMeASine-6c993962.js","assets/src_rust-d649cb4c.js","assets/GiveMeASine-d059f471.css"])}),"/magic_square":Z({asyncComponent:()=>J(()=>import("./Container-7db70681.js"),["assets/Container-7db70681.js","assets/src_rust-d649cb4c.js","assets/Container-f75cbe15.css"])})};return Be(()=>{let l=yt(localStorage.getItem("ns_site_section"));oe.update(s=>l)}),[o,r,i]}class Lt extends le{constructor(t){super(),ae(this,t,Ot,jt,F,{})}}new Lt({target:document.getElementById("app")});export{Ft as A,V as B,Gt as C,be as D,ye as E,Dt as F,Nt as G,Xt as H,At as I,zt as J,Bt as K,R as L,It as M,Tt as N,Be as O,Wt as P,Ce as Q,Ut as R,le as S,Vt as T,ce as U,ie as V,rt as W,Ht as X,et as Y,P as a,tt as b,y as c,E as d,C as e,Yt as f,Rt as g,me as h,ae as i,G as j,M as k,j as l,S as m,$ as n,it as o,A as p,_e as q,ze as r,F as s,Ae as t,k as u,q as v,x as w,z as x,oe as y,he as z};
