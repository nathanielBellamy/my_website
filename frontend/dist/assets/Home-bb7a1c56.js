import{S as qt,i as At,s as jt,e as l,a as Te,n as bt,d as Ue,b as Lt,o as Ot,c as r,f as Pe,t as Me,g as w,h as yt,L as j,j as p,k as m,l as b,m as e,p as $,q as Se,r as d,u as M,v as S,w as k,x as It,y as $t,I as Ht,z as Et,A as zt,B as Fe,C as Gt}from"./index-e5266db2.js";const Wt="/assets/magic_square_example-24e99fcd.gif",Ct="/assets/give_me_a_sine_example-7ccb4042.gif",Rt="/assets/0-980194f1.png",Bt="/assets/1-e3c74401.png",Dt="/assets/2-65d8bf2f.png",Ft="/assets/3-9b988824.png",Pt="/assets/4-3dc0b2b3.png",Tt="/assets/5-75498a8d.png",Ut="/assets/6-ac74c8e4.png",Vt="/assets/7-173f37c0.png",Jt="/assets/8-7b696f87.png",Kt="/assets/9-393e8585.png",Nt="/assets/10-3d8838a3.png",Qt="/assets/11-6a66497a.png",Xt="/assets/12-73e89b3d.png",Yt="/assets/13-4f7a9713.png",Zt="/assets/14-11ae09e9.png",xt="/assets/15-2d2a73e5.png",es="/assets/16-a6fdaa02.png",ts="/assets/17-b20e0717.png",ss="/assets/18-b005abc7.png",is="/assets/19-bc9e8afa.png",as="/assets/20-7542b5ea.png",ls="/assets/21-1ce06bfc.png",ns="/assets/22-4667002a.png",rs="/assets/23-e0a7f325.png",gs="/assets/24-1f5cdc26.png",os="/assets/25-1a413b78.png",fs="/assets/26-c2070eb6.png",ms="/assets/27-f9e00b01.png",cs="/assets/28-7e14bc2d.png",ps="/assets/29-c69d631e.png",us="/assets/30-a24987d9.png",_s="/assets/31-5b4d62ed.png",ds="/assets/32-ebfa12bc.png",vs="/assets/33-e7590edc.png",hs="/assets/34-9289b228.png",ws="/assets/35-f6748aca.png",bs="/assets/36-6e52604e.png",$s="/assets/37-ef7911d4.png",Ms="/assets/38-00adbff1.png",Ss="/assets/39-b27ed41e.png",ks="/assets/40-030be502.png",qs="/assets/41-772a95fe.png",As="/assets/42-00a7477b.png",js="/assets/43-803435fc.png",Ls="/assets/44-7585d9ed.png",Os="/assets/45-2b332a27.png",ys="/assets/46-7c265832.png";function Mt(t,i,o){const a=t.slice();return a[6]=i[o],a[8]=o,a}function St(t){let i,o;return{c(){i=l("img"),r(i,"class","h-full w-full ai_me ai_me_img svelte-14jqnve"),Pe(i.src,o=t[6])||r(i,"src",o),r(i,"alt",`AI ME #${t[8]}`),Me(i,"ai_me_img_hide",t[1]!==t[8]),Me(i,"ai_me_img_show",t[1]===t[8]),w(i,"height",t[0]),w(i,"width",t[0])},m(a,n){Te(a,i,n)},p(a,n){n&2&&Me(i,"ai_me_img_hide",a[1]!==a[8]),n&2&&Me(i,"ai_me_img_show",a[1]===a[8]),n&1&&w(i,"height",a[0]),n&1&&w(i,"width",a[0])},d(a){a&&Ue(i)}}}function Is(t){let i,o=t[2],a=[];for(let n=0;n<o.length;n+=1)a[n]=St(Mt(t,o,n));return{c(){i=l("div");for(let n=0;n<a.length;n+=1)a[n].c()},m(n,v){Te(n,i,v);for(let g=0;g<a.length;g+=1)a[g]&&a[g].m(i,null)},p(n,[v]){if(v&7){o=n[2];let g;for(g=0;g<o.length;g+=1){const _=Mt(n,o,g);a[g]?a[g].p(_,v):(a[g]=St(_),a[g].c(),a[g].m(i,null))}for(;g<a.length;g+=1)a[g].d(1);a.length=o.length}},i:bt,o:bt,d(n){n&&Ue(i),Lt(a,n)}}}function kt(t,i){return Math.floor(Math.random()*(i-t+1)+t)}function Hs(t,i,o){let a,{imgSideLength:n}=i;const v=[Rt,Bt,Dt,Ft,Pt,Tt,Ut,Vt,Jt,Kt,Nt,Qt,Xt,Yt,Zt,xt,es,ts,ss,is,as,ls,ns,rs,gs,os,fs,ms,cs,ps,us,_s,ds,vs,hs,ws,bs,$s,Ms,Ss,ks,qs,As,js,Ls,Os,ys];let g=kt(0,46);const h=setInterval(()=>o(3,g=kt(0,46)),5e3);return Ot(()=>{clearInterval(h)}),t.$$set=L=>{"imgSideLength"in L&&o(0,n=L.imgSideLength)},t.$$.update=()=>{t.$$.dirty&8&&o(1,a=g%47)},[n,a,v,g]}class Es extends qt{constructor(i){super(),At(this,i,Hs,Is,jt,{imgSideLength:0})}}function zs(t){let i,o,a,n=t[3].t("title",t[1])+"",v,g,_,h,L=t[3].t("intro/2",t[1])+"",Q,Y,u,E,Ve,R,Je,B,Ke,Z=t[3].t("intro/3",t[1])+"",ke,Ne,D,Qe,F,Xe,qe,x=t[3].t("intro/4",t[1])+"",Ae,Ye,ee,te=t[3].t("whatsHere",t[1])+"",je,Ze,se,H,O,ie,ae=t[3].t("about",t[1])+"",Le,xe,le,ne,z,et,re,P,Oe,ge=t[3].t("about_1",t[1])+"",ye,tt,Ie,oe=t[3].t("about_2",t[1])+"",He,st,y,fe,me=t[3].t("magicSquare",t[1])+"",Ee,it,ce,q,at,lt,pe,T,G,ue=t[3].t("magicSquare_1",t[1])+"",ze,nt,U,rt,V,gt,X,_e=t[3].t("magicSquare_2",t[1])+"",Ge,ot,W,ft,I,de,ve=t[3].t("giveMeASine",t[1])+"",We,mt,he,A,ct,pt,we,J,Ce,be=t[3].t("giveMeASine_1",t[1])+"",Re,ut,C,K,_t,$e=t[3].t("giveMeASine_2",t[1])+"",Be,dt,N,c,De,vt;return yt(t[5]),E=new j({props:{href:"https://www.rust-lang.org/",title:"Rust",sameOrigin:!1}}),R=new j({props:{href:"https://svelte.dev/",title:"Svelte",sameOrigin:!1}}),B=new j({props:{href:"https://www.typescriptlang.org/",title:"Typescript",sameOrigin:!1}}),D=new j({props:{href:"https://webassembly.org/",title:"WebAssembly",sameOrigin:!1}}),F=new j({props:{href:"https://crates.io/crates/wasm-bindgen",title:"wasm-bindgen.",sameOrigin:!1}}),z=new Es({props:{imgSideLength:t[2]}}),U=new j({props:{href:"https://www.khronos.org/webgl/",title:"WebGL",sameOrigin:!1}}),V=new j({props:{href:"https://rustwasm.github.io/wasm-bindgen/examples/webgl.html",title:"RustWasm",sameOrigin:!1}}),W=new j({props:{href:"https://en.wikipedia.org/wiki/Modular_synthesizer",title:t[3].t("magicSquare_3",t[1]),sameOrigin:!1}}),K=new j({props:{href:"https://rustwasm.github.io/wasm-bindgen",title:"RustWasm",sameOrigin:!1}}),N=new j({props:{href:"/magic_square",title:"Magic Square",sameOrigin:!0}}),{c(){i=l("body"),o=l("div"),a=l("h2"),v=p(n),g=m(),_=l("ul"),h=l("li"),Q=p(L),Y=m(),u=l("p"),b(E.$$.fragment),Ve=p(`
          +
          `),b(R.$$.fragment),Je=p(`
          +
          `),b(B.$$.fragment),Ke=m(),ke=p(Z),Ne=m(),b(D.$$.fragment),Qe=p(`
        +
        `),b(F.$$.fragment),Xe=m(),qe=l("li"),Ae=p(x),Ye=m(),ee=l("div"),je=p(te),Ze=m(),se=l("div"),H=l("div"),O=l("button"),ie=l("div"),Le=p(ae),xe=m(),le=l("div"),ne=l("div"),b(z.$$.fragment),et=m(),re=l("div"),P=l("ul"),Oe=l("li"),ye=p(ge),tt=m(),Ie=l("li"),He=p(oe),st=m(),y=l("button"),fe=l("div"),Ee=p(me),it=m(),ce=l("div"),q=l("img"),lt=m(),pe=l("div"),T=l("ul"),G=l("li"),ze=p(ue),nt=m(),b(U.$$.fragment),rt=p(`
              +
              `),b(V.$$.fragment),gt=m(),X=l("li"),Ge=p(_e),ot=m(),b(W.$$.fragment),ft=m(),I=l("button"),de=l("div"),We=p(ve),mt=m(),he=l("div"),A=l("img"),pt=m(),we=l("div"),J=l("ul"),Ce=l("li"),Re=p(be),ut=m(),C=l("li"),b(K.$$.fragment),_t=m(),Be=p($e),dt=m(),b(N.$$.fragment),r(a,"class","home_title text-left pl-5 svelte-lc9rxf"),r(_,"class","home_intro_list text-left p-5 flex flex-col justify-between items-stretch svelte-lc9rxf"),r(o,"class","home_title_container flex flex-col justify-between items-stretch md:flex-row md:justify-start md:items-center svelte-lc9rxf"),r(ee,"class","home_title_dark text-left pl-5 svelte-lc9rxf"),r(ie,"class","preview_title svelte-lc9rxf"),r(ne,"class","ai_me_container magic_square_img grid grid-rows-1 grid-cols-1 svelte-lc9rxf"),r(le,"class","row-span-2 w-full flex justify-around items-center"),r(P,"class","preview_list svelte-lc9rxf"),r(re,"class","flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll"),r(O,"class","preview grid grid-cols-1 grid-rows-4 gap-2 svelte-lc9rxf"),r(fe,"class","preview_title svelte-lc9rxf"),r(q,"class","magic_square_img ai_me svelte-lc9rxf"),Pe(q.src,at=Wt)||r(q,"src",at),r(q,"alt","Magic Square Example"),w(q,"height",t[2]),w(q,"width",t[2]),r(ce,"class","row-span-2 flex justify-around items-center"),r(T,"class","preview_list svelte-lc9rxf"),r(pe,"class","flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll"),r(y,"class","preview grid grid-cols-1 grid-rows-4 gap-2 svelte-lc9rxf"),r(de,"class","preview_title svelte-lc9rxf"),r(A,"class","magic_square_img svelte-lc9rxf"),Pe(A.src,ct=Ct)||r(A,"src",ct),r(A,"alt","Give Me A Sine Example"),w(A,"height",t[2]),w(A,"width",t[2]),r(he,"class","row-span-2 flex justify-around items-center"),r(J,"class","preview_list svelte-lc9rxf"),r(we,"class","flex flex-col pl-5 pr-5 mb-2 justify-around items-stretch overflow-y-scroll"),r(I,"class","preview grid grid-cols-1 grid-rows-4 gap-2 svelte-lc9rxf"),r(H,"class","grow grid grid-cols-1 grid-rows-3 md:grid-cols-3 md:grid-rows-1 gap-3"),r(se,"class","grow pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch"),r(i,"class","pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch gap-2")},m(s,f){Te(s,i,f),e(i,o),e(o,a),e(a,v),e(o,g),e(o,_),e(_,h),e(h,Q),e(h,Y),e(h,u),$(E,u,null),e(u,Ve),$(R,u,null),e(u,Je),$(B,u,null),e(u,Ke),e(u,ke),e(h,Ne),$(D,h,null),e(h,Qe),$(F,h,null),e(_,Xe),e(_,qe),e(qe,Ae),e(i,Ye),e(i,ee),e(ee,je),e(i,Ze),e(i,se),e(se,H),e(H,O),e(O,ie),e(ie,Le),e(O,xe),e(O,le),e(le,ne),$(z,ne,null),e(O,et),e(O,re),e(re,P),e(P,Oe),e(Oe,ye),e(P,tt),e(P,Ie),e(Ie,He),e(H,st),e(H,y),e(y,fe),e(fe,Ee),e(y,it),e(y,ce),e(ce,q),e(y,lt),e(y,pe),e(pe,T),e(T,G),e(G,ze),e(G,nt),$(U,G,null),e(G,rt),$(V,G,null),e(T,gt),e(T,X),e(X,Ge),e(X,ot),$(W,X,null),e(H,ft),e(H,I),e(I,de),e(de,We),e(I,mt),e(I,he),e(he,A),e(I,pt),e(I,we),e(we,J),e(J,Ce),e(Ce,Re),e(J,ut),e(J,C),$(K,C,null),e(C,_t),e(C,Be),e(C,dt),$(N,C,null),c=!0,De||(vt=[Se(window,"resize",t[5]),Se(O,"click",t[6]),Se(y,"click",t[7]),Se(I,"click",t[8])],De=!0)},p(s,[f]){(!c||f&2)&&n!==(n=s[3].t("title",s[1])+"")&&d(v,n),(!c||f&2)&&L!==(L=s[3].t("intro/2",s[1])+"")&&d(Q,L),(!c||f&2)&&Z!==(Z=s[3].t("intro/3",s[1])+"")&&d(ke,Z),(!c||f&2)&&x!==(x=s[3].t("intro/4",s[1])+"")&&d(Ae,x),(!c||f&2)&&te!==(te=s[3].t("whatsHere",s[1])+"")&&d(je,te),(!c||f&2)&&ae!==(ae=s[3].t("about",s[1])+"")&&d(Le,ae);const ht={};f&4&&(ht.imgSideLength=s[2]),z.$set(ht),(!c||f&2)&&ge!==(ge=s[3].t("about_1",s[1])+"")&&d(ye,ge),(!c||f&2)&&oe!==(oe=s[3].t("about_2",s[1])+"")&&d(He,oe),(!c||f&2)&&me!==(me=s[3].t("magicSquare",s[1])+"")&&d(Ee,me),f&4&&w(q,"height",s[2]),f&4&&w(q,"width",s[2]),(!c||f&2)&&ue!==(ue=s[3].t("magicSquare_1",s[1])+"")&&d(ze,ue),(!c||f&2)&&_e!==(_e=s[3].t("magicSquare_2",s[1])+"")&&d(Ge,_e);const wt={};f&2&&(wt.title=s[3].t("magicSquare_3",s[1])),W.$set(wt),(!c||f&2)&&ve!==(ve=s[3].t("giveMeASine",s[1])+"")&&d(We,ve),f&4&&w(A,"height",s[2]),f&4&&w(A,"width",s[2]),(!c||f&2)&&be!==(be=s[3].t("giveMeASine_1",s[1])+"")&&d(Re,be),(!c||f&2)&&$e!==($e=s[3].t("giveMeASine_2",s[1])+"")&&d(Be,$e)},i(s){c||(M(E.$$.fragment,s),M(R.$$.fragment,s),M(B.$$.fragment,s),M(D.$$.fragment,s),M(F.$$.fragment,s),M(z.$$.fragment,s),M(U.$$.fragment,s),M(V.$$.fragment,s),M(W.$$.fragment,s),M(K.$$.fragment,s),M(N.$$.fragment,s),c=!0)},o(s){S(E.$$.fragment,s),S(R.$$.fragment,s),S(B.$$.fragment,s),S(D.$$.fragment,s),S(F.$$.fragment,s),S(z.$$.fragment,s),S(U.$$.fragment,s),S(V.$$.fragment,s),S(W.$$.fragment,s),S(K.$$.fragment,s),S(N.$$.fragment,s),c=!1},d(s){s&&Ue(i),k(E),k(R),k(B),k(D),k(F),k(z),k(U),k(V),k(W),k(K),k(N),De=!1,It(vt)}}}function Gs(t){return Math.floor(t/3.5).toString()+"px"}function Ws(t,i,o){let a;$t.subscribe(u=>u);const n=new Ht("home");let v;Et.subscribe(u=>o(1,v=u));let g;function _(u){localStorage.setItem("ns_site_section",u),$t.update(E=>u),zt(Gt(u))}function h(){o(0,g=window.innerHeight)}const L=()=>_(Fe.about),Q=()=>_(Fe.magicSquare),Y=()=>_(Fe.giveMeASine);return t.$$.update=()=>{t.$$.dirty&1&&o(2,a=Gs(g))},[g,v,a,n,_,h,L,Q,Y]}class Rs extends qt{constructor(i){super(),At(this,i,Ws,zs,jt,{})}}export{Rs as default};
