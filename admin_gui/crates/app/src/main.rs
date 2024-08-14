use gpui::*;
use ui::helloworld::HelloWorld;

fn main() {
    App::new().run(|cx: &mut AppContext| {
        let bounds = Bounds::centered(None, size(px(600.0), px(600.0)), cx);
        let _ = cx.open_window(
            WindowOptions {
                window_bounds: Some(WindowBounds::Windowed(bounds)),
                ..Default::default()
            },
            |cx| {
                cx.new_view(|_cx| HelloWorld {
                    text: "Hello World !".into(),
                })
            },
        );
    });
}
