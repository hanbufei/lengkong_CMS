use iced::Sandbox;

pub struct Helloworld;

impl Sandbox for Helloworld {
    type Message = ();

    fn new() -> Self {
        Self
    }

    fn title(&self) -> String {
        String::from("My App")
    }

    fn update(&mut self, _message: Self::Message) {}

    fn view(&self) -> iced::Element<Self::Message> {
        "Hello World!".into()
    }
}
