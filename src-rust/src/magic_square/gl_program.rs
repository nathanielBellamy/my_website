use web_sys::{WebGl2RenderingContext, WebGlProgram};
use super::gl_shader::GlShader;

pub struct GlProgram;

impl GlProgram {
    pub fn new(gl: &WebGl2RenderingContext) -> Result<WebGlProgram, String> {
        let gl_program = gl
            .create_program()
            .ok_or_else(|| format!("Unable to create gl_program"))?;
        
        gl.attach_shader(
            &gl_program, 
            &GlShader::vert(gl)?
        );
        gl.attach_shader(
            &gl_program, 
            &GlShader::frag(gl)?
        );
        gl.link_program(&gl_program);

        if gl
            .get_program_parameter(&gl_program, WebGl2RenderingContext::LINK_STATUS)
            .as_bool()
            .unwrap_or(false)
        {
            Ok(gl_program)
        } else {
            Err(gl
                .get_program_info_log(&gl_program)
                .unwrap_or_else(|| String::from("Unknown error creating program object")))
        }
    }
}

