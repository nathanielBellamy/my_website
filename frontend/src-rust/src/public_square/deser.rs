pub struct Deser;

impl Deser {
    // let bytes: &[u8] = unsafe { any_as_u8_slice(&my_struct) };
    pub unsafe fn any_as_u8_slice<T: Sized>(p: &T) -> &[u8] {
        ::core::slice::from_raw_parts(
            (p as *const T) as *const u8,
            ::core::mem::size_of::<T>(),
        )
    }
}

