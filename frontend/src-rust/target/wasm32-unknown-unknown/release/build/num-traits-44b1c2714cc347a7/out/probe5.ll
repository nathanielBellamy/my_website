; ModuleID = 'probe5.8a55f6231c30058c-cgu.0'
source_filename = "probe5.8a55f6231c30058c-cgu.0"
target datalayout = "e-m:e-p:32:32-p10:8:8-p20:8:8-i64:64-n32:64-S128-ni:1:10:20"
target triple = "wasm32-unknown-unknown"

@alloc_81650af4d9bc8ae635c155fc21cc72ff = private unnamed_addr constant <{ [75 x i8] }> <{ [75 x i8] c"/rustc/2f6bc5d259e7ab25ddfdd33de53b892770218918/library/core/src/num/mod.rs" }>, align 1
@alloc_26b935afffb7853b7bf8683b93488d14 = private unnamed_addr constant <{ ptr, [12 x i8] }> <{ ptr @alloc_81650af4d9bc8ae635c155fc21cc72ff, [12 x i8] c"K\00\00\00I\04\00\00\05\00\00\00" }>, align 4
@str.0 = internal constant [25 x i8] c"attempt to divide by zero"

; probe5::probe
; Function Attrs: nounwind
define hidden void @_ZN6probe55probe17ha6e06302acffe969E() unnamed_addr #0 {
start:
  %0 = call i1 @llvm.expect.i1(i1 false, i1 false)
  br i1 %0, label %panic.i, label %"_ZN4core3num21_$LT$impl$u20$u32$GT$10div_euclid17h4dbf0abbbe944c34E.exit"

panic.i:                                          ; preds = %start
; call core::panicking::panic
  call void @_ZN4core9panicking5panic17h3cf1911af20028feE(ptr align 1 @str.0, i32 25, ptr align 4 @alloc_26b935afffb7853b7bf8683b93488d14) #3
  unreachable

"_ZN4core3num21_$LT$impl$u20$u32$GT$10div_euclid17h4dbf0abbbe944c34E.exit": ; preds = %start
  ret void
}

; Function Attrs: nocallback nofree nosync nounwind willreturn memory(none)
declare hidden i1 @llvm.expect.i1(i1, i1) #1

; core::panicking::panic
; Function Attrs: cold noinline noreturn nounwind
declare dso_local void @_ZN4core9panicking5panic17h3cf1911af20028feE(ptr align 1, i32, ptr align 4) unnamed_addr #2

attributes #0 = { nounwind "target-cpu"="generic" }
attributes #1 = { nocallback nofree nosync nounwind willreturn memory(none) }
attributes #2 = { cold noinline noreturn nounwind "target-cpu"="generic" }
attributes #3 = { noreturn nounwind }
