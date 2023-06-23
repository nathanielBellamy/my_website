use web_sys::{WebGl2RenderingContext, WebGlShader};
use super::main::log;

pub struct GlShader;

impl GlShader {
    fn exec(
        gl: &WebGl2RenderingContext,
        shader_type: u32,
        source: &str,
    ) -> Result<WebGlShader, String> {
        let gl_shader = gl
            .create_shader(shader_type)
            .ok_or_else(|| String::from("Unable to create shader object"))?;
        gl.shader_source(&gl_shader, source);
        gl.compile_shader(&gl_shader);

        if gl
            .get_shader_parameter(&gl_shader, WebGl2RenderingContext::COMPILE_STATUS)
            .as_bool()
            .unwrap_or(false)
        {
            Ok(gl_shader)
        } else {
            let err = gl
                .get_shader_info_log(&gl_shader)
                .unwrap_or_else(|| String::from("Unknown error creating shader"));

            log(&err);
            Err(err)
        }
    }

    pub fn frag(gl: &WebGl2RenderingContext) -> Result<WebGlShader, String> {
        GlShader::exec(
            gl, 
            WebGl2RenderingContext::FRAGMENT_SHADER, 
            r##"#version 300 es
            precision highp float;
            out vec4 outColor;

            uniform vec4 u_rgba;
            
            void main() {{
                outColor = u_rgba;
            }}
            "##,
        )
    }

    pub fn vert(gl: &WebGl2RenderingContext) -> Result<WebGlShader, String>{
        GlShader::exec(
            &gl,
            WebGl2RenderingContext::VERTEX_SHADER,
            r##"#version 300 es
 
            in vec4 position;

            uniform vec4 u_translation;

            void main() {
                gl_Position = u_translation + position;
            }
            "##,
        )
    }
}

