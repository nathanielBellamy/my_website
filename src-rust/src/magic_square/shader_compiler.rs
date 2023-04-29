
use web_sys::{WebGl2RenderingContext, WebGlShader};
use crate::magic_square::main::Rgba;

pub struct ShaderCompiler;

impl ShaderCompiler {
    pub fn exec(
        context: &WebGl2RenderingContext,
        shader_type: u32,
        source: &str,
    ) -> Result<WebGlShader, String> {
        let shader = context
            .create_shader(shader_type)
            .ok_or_else(|| String::from("Unable to create shader object"))?;
        context.shader_source(&shader, source);
        context.compile_shader(&shader);

        if context
            .get_shader_parameter(&shader, WebGl2RenderingContext::COMPILE_STATUS)
            .as_bool()
            .unwrap_or(false)
        {
            Ok(shader)
        } else {
            Err(context
                .get_shader_info_log(&shader)
                .unwrap_or_else(|| String::from("Unknown error creating shader")))
        }
    }

    pub fn frag_default(
        context: &WebGl2RenderingContext,
        rgba: &Rgba,
    ) -> Result<WebGlShader, String> {
        let string = format!(
            r##"#version 300 es
                precision highp float;
                out vec4 outColor;
                
                void main() {{
                    outColor = vec4({}, {}, {}, {});
                }}
                "##,
            rgba[0], rgba[1], rgba[2], rgba[3]
        );

        ShaderCompiler::exec(&context, WebGl2RenderingContext::FRAGMENT_SHADER, &string)
    }

    pub fn vert_default(context: &WebGl2RenderingContext) -> Result<WebGlShader, String> {
        ShaderCompiler::exec(
            &context,
            WebGl2RenderingContext::VERTEX_SHADER,
            r##"#version 300 es
     
            in vec4 position;

            void main() {
            
                gl_Position = position;
            }
            "##,
        )
    }
}

