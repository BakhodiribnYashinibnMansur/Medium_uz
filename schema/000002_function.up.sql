-- //////////////////////////////////////////////////////////////////////////
-- LIKE TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION like_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_like_count = post_like_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER like_post_trigger AFTER INSERT ON liked_post
FOR EACH ROW EXECUTE PROCEDURE like_post();

-- //////////////////////////////////////////////////////////////////////////
-- UNLIKE TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION unlike_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_like_count = post_like_count - 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER unlike_post_trigger AFTER UPDATE ON liked_post
FOR EACH ROW EXECUTE PROCEDURE unlike_post();


-- //////////////////////////////////////////////////////////////////////////
-- VIEW TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION view_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_views_count = post_views_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER view_post_trigger AFTER INSERT ON viewed_posts
FOR EACH ROW EXECUTE PROCEDURE view_post();
