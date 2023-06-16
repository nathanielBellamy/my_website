use std::rc::Rc;
use std::cell::RefCell;
use super::animation::Animation;
use super::main::Axis;
use super::main::log;
use super::transformations::{RotationSequence, Rotation, Transformation, Translation};
use super::ui_buffer::UiBuffer;
use super::geometry::cache::{Cache as GeometryCache, CACHE_CAPACITY};
use super::geometry::{Geometry, Shape};
use super::settings::MouseTracking;
use super::vertices::VERTEX_ARRAY_SIZE;


pub struct Render;

impl Render {
    pub fn all_lines(
        mouse_pos_buffer: &Rc<RefCell<[f32; 2]>>,
        ui_buffer: &UiBuffer,
        // geometry_cache: Arc<Mutex<GeometryCache>>,
        geometry_cache: &Rc<RefCell<GeometryCache>>,
        // animation: &Rc<RefCell<Animation>>,
    ) {
        // let max_idx = Settings::max_idx_from_draw_pattern(ui_buffer.settings.draw_pattern);
        let mouse_pos_buffer = *mouse_pos_buffer.clone().borrow();
        let animation = Animation::new_from_settings(&ui_buffer.settings);

        for idx in 0..CACHE_CAPACITY { // geometry_cache.max_idx + 1 { //TODO: settings.cache_per
            let idx_f32 = idx as f32;
            let rot_seq = RotationSequence::new(
                Rotation::new(
                    Axis::X, 
                    ui_buffer.settings.x_rot_base
                        + (mouse_pos_buffer[0] 
                            + ui_buffer.settings.translation_x_base) * ui_buffer.settings.x_axis_x_rot_coeff
                        + (mouse_pos_buffer[1] 
                            + ui_buffer.settings.translation_y_base) * ui_buffer.settings.y_axis_x_rot_coeff
                        + idx_f32 * ui_buffer.settings.x_rot_spread
                ),
                Rotation::new(
                    Axis::Y,
                    ui_buffer.settings.y_rot_base
                        + (mouse_pos_buffer[0] 
                            + ui_buffer.settings.translation_x_base) * ui_buffer.settings.x_axis_y_rot_coeff
                        + (mouse_pos_buffer[1] 
                            + ui_buffer.settings.translation_y_base) * ui_buffer.settings.y_axis_y_rot_coeff
                        + idx_f32 * ui_buffer.settings.y_rot_spread
                ),
                Rotation::new(
                    Axis::Z,
                    ui_buffer.settings.z_rot_base
                        + (mouse_pos_buffer[0] 
                            + ui_buffer.settings.translation_x_base) * ui_buffer.settings.x_axis_z_rot_coeff
                        + (mouse_pos_buffer[1] 
                            + ui_buffer.settings.translation_y_base) * ui_buffer.settings.y_axis_z_rot_coeff
                        + idx_f32 * ui_buffer.settings.z_rot_spread
                ),
            );

            let translation = match ui_buffer.settings.mouse_tracking {
                MouseTracking::On => Translation { 
                    x: ui_buffer.settings.translation_x_base
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base 
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base 
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::Off => Translation { 
                    x: ui_buffer.settings.translation_x_base
                        + (idx_f32 * ui_buffer.settings.translation_x_spread), 
                    y: ui_buffer.settings.translation_y_base
                        - (idx_f32 * ui_buffer.settings.translation_y_spread), 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::InvX =>  Translation { 
                    x: ui_buffer.settings.translation_x_base 
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::InvY =>  Translation { 
                    x: ui_buffer.settings.translation_x_base 
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base 
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::InvXY =>  Translation { 
                    x: ui_buffer.settings.translation_x_base
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
            };
            
            // TODO: Debug animation.set_real()
            let mut shape: Shape = animation.reel[animation.idx][idx];
        
            // set for testing
            // proof of concept for refactor
            shape = Shape::Icosahedron;
            if idx % 2 == 1 {
                shape = Shape::Hexagon;
            }

            log(&format!("{:?}", shape));

            let radius: f32 = ui_buffer.settings.radius_step * idx_f32 + ui_buffer.settings.radius_base; 
            let geometry: Geometry = Geometry {
                radius,
                transformation: Transformation {
                    order: ui_buffer.settings.transform_order,
                    rot_seq,
                    translation
                }
            };
            
            // log(&format!("geogeogeo {:?}", geometry));
            let shape_vertex_arr: [f32; VERTEX_ARRAY_SIZE] = match shape {
                Shape::None => geometry.empty(), 
                Shape::Icosahedron => geometry.icosahedron(),
                Shape::Hexagon => geometry.hexagon(),
                _ => geometry.icosahedron(),
            };
            
            // log(&format!("arr: {:?}", shape_vertex_arr));

            geometry_cache.borrow_mut().set_next(shape_vertex_arr, shape);
            // log(&format!("geo_cache {:?}", geometry_cache));
        }
    }
}
