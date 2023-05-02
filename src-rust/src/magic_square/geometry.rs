
pub struct Geometry;

impl Geometry {
    pub fn hexagon(buffer: &[f32; 2], radius: f32, rotation: RotationSequence) -> Vertices {
        let mut vertices = Vertices::new();

        let center_x = buffer[0];
        let center_y = buffer[1];
        let xy = 0.02 * center_x * center_y;

        let x_shift = radius * 0.5; // r cos(pi/3)
        let y_shift = radius * 0.86602540378; // r sin(pi/3)



        // draw hexagon boundary
        // start north east corner
        // end east corner
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south east corner
        // end south west corner
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south west corner
        // end west corner
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );

        // start west corner
        // end north west corner
        vertices.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );


        // start north west corner
        // end north east corner
        vertices.set_next( 
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );

        vertices
    }
}


