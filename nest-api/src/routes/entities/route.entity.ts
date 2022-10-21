import { Prop, Schema, raw, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type RouteDocument = Route & Document;

@Schema()
export class Route {
  @Prop()
  _id: string;

  @Prop()
  title: string;

  @Prop(
    // tipo do campo na collection no mongodb
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  startPosition: { lat: number; lng: number };

  @Prop(
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  endPosition: { lat: number; lng: number };
}

export const RouteSchema = SchemaFactory.createForClass(Route);
